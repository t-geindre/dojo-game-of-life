package d05

import (
	"math/rand"
	"runtime"
	"sync"
)

type Grid struct {
	cells, next []bool
	w, h        int
}

func NewGrid(w, h int) *Grid {
	return &Grid{
		cells: make([]bool, w*h),
		next:  make([]bool, w*h),
		w:     w,
		h:     h,
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
	numCores := runtime.NumCPU()
	chunkSize := (len(g.cells) + numCores - 1) / numCores
	var wg sync.WaitGroup

	for i := 0; i < numCores; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(g.cells) {
			end = len(g.cells)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				n := g.CountAliveNeighbors(j)
				if g.cells[j] {
					g.next[j] = n == 2 || n == 3
				} else {
					g.next[j] = n == 3
				}
			}
		}(start, end)
	}

	wg.Wait()
	g.cells, g.next = g.next, g.cells
}
