package d10

import (
	"dojo-game-of-life/controls"
	"dojo-game-of-life/debug"
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Game struct {
	pixels []byte
	grid   *Grid
	buff   *ebiten.Image

	zoom  *controls.Wheel
	speed *controls.Wheel
	move  *controls.Move

	scale  float64
	tx, ty float64

	skip    float64
	skipped int
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)
	grid.Randomize(.3)

	speed := controls.NewWheel()
	speed.Min = 0
	speed.Max = 100
	speed.Modifier = ebiten.KeyAlt
	speed.Cursor = ebiten.CursorShapeEWResize
	speed.Rate = 1
	speed.Speed = 1
	speed.SetValue(100)

	return &Game{
		pixels: make([]byte, w*h*4),
		grid:   grid,
		buff:   ebiten.NewImage(w, h),
		zoom:   controls.NewWheel(),
		speed:  speed,
		move:   controls.NewMove(),
		scale:  1,
	}
}

func (g *Game) Update() {
	g.scale = g.zoom.Value()
	g.skip = g.speed.Value()

	fx, fy := g.zoom.Focus()
	g.tx, g.ty = g.move.ValueWithDeltaScaleAt(g.zoom.Delta(), fx, fy)

	if g.skipped++; float64(g.skipped) > 100-g.skip && g.skip > 0 {
		g.skipped = 0
		g.grid.NextState()
		g.UpdatePixels()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.buff.WritePixels(g.pixels)

	srcBds := g.buff.Bounds()
	dstBds := screen.Bounds()

	sw := float64(srcBds.Dx()) * g.scale
	sh := float64(srcBds.Dy()) * g.scale

	dw := float64(dstBds.Dx())
	dh := float64(dstBds.Dy())

	ox := math.Mod(-g.tx, sw)
	oy := math.Mod(-g.ty, sh)

	if ox > 0 {
		ox -= sw
	}
	if oy > 0 {
		oy -= sh
	}

	n := 0
	for x := ox; x < dw; x += sw {
		for y := oy; y < dh; y += sh {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Scale(g.scale, g.scale)
			opt.Filter = ebiten.FilterNearest
			opt.GeoM.Translate(x, y)
			screen.DrawImage(g.buff, opt)
			n++
		}
	}

	debug.DrawPrintf(
		screen, debug.BottomRight,
		"[CTL]+[MW] Scale (%0.f%%)\n[ALT]+[MW] Speed (%0.f%%)\n[SPACE]+[LMB] Move",
		g.scale*100, g.skip,
	)

	debug.DrawPrintf(screen, debug.BottomLeft, "Repetitions %d, dz %.2f", n, g.zoom.Delta())
}

func (g *Game) UpdatePixels() {
	for i, c := range g.grid.Cells() {
		v := byte(0)
		if c {
			v = 0xff
		}
		g.pixels[i*4] = v
		g.pixels[i*4+1] = v
		g.pixels[i*4+2] = v
		g.pixels[i*4+3] = 0xff
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
