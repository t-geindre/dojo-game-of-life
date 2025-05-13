package d01

import "math/rand"

type Grid struct {
	cells []bool
	w, h  int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		w: w, h: h,
		// Even if the grid is two-dimensional, we can store it as a one-dimensional array
		// This is way more efficient
		cells: make([]bool, w*h),
	}
}

func (g *Grid) Cells() []bool {
	return g.cells
}

func (g *Grid) Size() (int, int) {
	return g.w, g.h
}

func (g *Grid) Randomize(f float64) {
	// Randomize the grid according to the given factor
	for i := range g.cells {
		g.cells[i] = rand.Float64() < f
	}
}
