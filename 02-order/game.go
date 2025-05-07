package d02

import (
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	pixel []byte
	grid  *Grid
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)

	// The grid is now randomized once
	grid.Randomize(.3)

	return &Game{
		pixel: make([]byte, w*h*4),
		grid:  grid,
	}
}

func (g *Game) Update() {
	// Now we want to compute the grid next state on each tick
	g.grid.NextState()

	// Update the pixels buffer to reflect the new grid state
	g.UpdatePixels()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixel)
}

func (g *Game) UpdatePixels() {
	for i, c := range g.grid.Cells() {
		v := byte(0)
		if c {
			v = 0xff
		}
		g.pixel[i*4] = v
		g.pixel[i*4+1] = v
		g.pixel[i*4+2] = v
		g.pixel[i*4+3] = 0xff
	}
}
