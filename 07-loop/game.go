package d07

import (
	"dojo-game-of-life/controls"
	"dojo-game-of-life/debug"
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	pixels []byte
	grid   *Grid
	buff   *ebiten.Image

	zoom  *controls.Wheel
	speed *controls.Wheel

	scale   float64
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
		scale:  1,
	}
}

func (g *Game) Update() {
	g.scale = g.zoom.Value()
	g.skip = g.speed.Value()

	if g.skipped++; float64(g.skipped) > 100-g.skip && g.skip > 0 {
		g.skipped = 0
		g.grid.NextState()
		g.UpdatePixels()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.buff.WritePixels(g.pixels)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(g.scale, g.scale)
	opt.Filter = ebiten.FilterNearest

	screen.DrawImage(g.buff, opt)

	debug.DrawPrintf(
		screen, debug.BottomRight,
		"[CTL]+[MW] Scale (%0.f%%)\n[ALT]+[MW] Speed (%0.f%%)",
		g.scale*100, g.skip,
	)
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
