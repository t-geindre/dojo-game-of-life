package d03

import (
	"dojo-game-of-life/debug"
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
	"runtime"
)

type Game struct {
	pixels []byte
	grid   *Grid
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)
	grid.Randomize(.3)

	return &Game{
		pixels: make([]byte, w*h*4),
		grid:   grid,
	}
}

func (g *Game) Update() {
	g.grid.NextState()
	g.UpdatePixels()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
	debug.DrawPrintf(screen, debug.TopRight, "%d CPU", runtime.NumCPU())
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
