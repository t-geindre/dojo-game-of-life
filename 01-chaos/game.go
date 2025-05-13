package d01

import (
	"dojo-game-of-life/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	pixels []byte
	grid   *Grid
}

func NewGame(w, h int) game.Game {
	return &Game{
		// A pixel is made of 4 bytes (Red, Green, Blue, Alpha)
		// Width * Height gives the total number of pixels
		pixels: make([]byte, w*h*4),

		// The grid is the object that will hold the game state
		grid: NewGrid(w, h),
	}
}

func (g *Game) Update() {
	// Randomize the grid with a 50% chance of being alive
	// The chance of being alive is parameter float64 between 0 and 1, 0.5 is 50%
	// ...
	g.grid.Randomize(0.5)

	// Update the game pixels
	// ...
	// Remember, a pixel is made of 4 bytes (Red, Green, Blue, Alpha)
	// Where the grid is only made of booleans
	// ...
	for i, c := range g.grid.Cells() {
		b := byte(0)
		if c {
			b = byte(255)
		}
		for j := 0; j < 4; j++ {
			g.pixels[i*4+j] = b
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the pixels to the screen
	screen.WritePixels(g.pixels) //should do the trick if pixels are in the right format
}
