package d05

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
	zoom   *controls.Wheel
	scale  float64
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)
	grid.Randomize(.3)

	return &Game{
		pixels: make([]byte, w*h*4),
		grid:   grid,
		buff:   ebiten.NewImage(w, h),
		zoom:   controls.NewWheel(),
		scale:  1,
	}
}

func (g *Game) Update() {
	g.scale = g.zoom.Value()
	g.grid.NextState()
	g.UpdatePixels()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.buff.WritePixels(g.pixels)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(g.scale, g.scale)
	opt.Filter = ebiten.FilterNearest

	screen.DrawImage(g.buff, opt)

	debug.DrawPrintf(
		screen, debug.BottomRight,
		"[CTL]+[MW] Zoom (%0.f%%)",
		g.scale*100,
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
