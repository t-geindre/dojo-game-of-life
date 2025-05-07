package d02

import "math/rand"

type Grid struct {
	cells, next []bool
	w, h        int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		w: w, h: h,
		cells: make([]bool, w*h),

		// Here is a second buffer to store the next state
		// This is because we need to compute the next state based on the current state
		// We can't just update the current state in place because it would affect the computation
		next: make([]bool, w*h),
	}
}

func (g *Grid) Cells() []bool {
	return g.cells
}

func (g *Grid) Size() (int, int) {
	return g.w, g.h
}

func (g *Grid) Randomize(f float64) {
	for i := range g.cells {
		g.cells[i] = rand.Float64() < f
	}
}

func (g *Grid) NextState() {
	for i := range g.cells {
		// Count the number of alive neighbors
		// and update the next state based on the rules of the game
		n := g.CountAliveNeighbors(i)

		// Reminder:
		// - Alive cell: stays alive if 2 or 3 neighbors, dies otherwise
		// - Dead cell: becomes alive if exactly 3 neighbors, stays dead otherwise
		// ...
		_ = n
	}

	// Swap the cells and next buffers
	g.cells, g.next = g.next, g.cells
}

func (g *Grid) CountAliveNeighbors(idx int) int {
	// We need to compute the number of alive neighbors for the given index
	// ...

	// We are working with to coordinates systems:
	// - The one-dimensional array index
	// - The two-dimensional grid coordinates
	//
	// To check all the neighbors, we need to check the 8 surrounding cells
	// which are in the grid coordinates system:
	// {-1, -1}, {0, -1}, {1, -1},
	// {-1, 0},           {1, 0},
	// {-1, 1},  {0, 1},  {1, 1},
	// ...

	// The first step is to convert the one-dimensional index to the two-dimensional coordinates
	// ...
	x, y := 0, 0 // modulo, division...

	// From there, we can establish the 8 surrounding coordinates
	// ...
	dirs := []struct{ dx, dy int }{
		{x, y}, // Fix me of course
	}

	// We start with 0 neighbors alive
	// ...
	n := 0

	for _, d := range dirs {
		// If the cell at the given coordinates is alive, we increment the counter
		// We must also check that the coordinates are within the grid bounds
		_ = d
	}

	return n
}
