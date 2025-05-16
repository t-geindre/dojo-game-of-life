package d01s

import "math/rand"

type Grid struct {
	cells []bool
	w, h  int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		w: w, h: h,
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
	for i := range g.cells {
		g.cells[i] = rand.Float64() < f
	}
}
