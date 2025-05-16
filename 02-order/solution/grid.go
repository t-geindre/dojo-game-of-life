package d02s

import "math/rand"

type Grid struct {
	cells, next []bool
	w, h        int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		w: w, h: h,
		cells: make([]bool, w*h),
		next:  make([]bool, w*h),
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

func (g *Grid) CountAliveNeighbors(idx int) int {
	x, y := idx%g.w, idx/g.w
	dirs := []struct{ dx, dy int }{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y}, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}
	n := 0
	for _, d := range dirs {
		if d.dx >= 0 && d.dx < g.w && d.dy >= 0 && d.dy < g.h {
			if g.cells[d.dy*g.w+d.dx] {
				n++
			}
		}
	}
	return n
}

func (g *Grid) NextState() {
	for i := range g.cells {
		n := g.CountAliveNeighbors(i)
		if g.cells[i] {
			g.next[i] = n == 2 || n == 3
		} else {
			g.next[i] = n == 3
		}
	}

	g.cells, g.next = g.next, g.cells
}
