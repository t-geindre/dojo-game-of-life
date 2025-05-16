package d05

import "github.com/hajimehoshi/ebiten/v2"

type GridRenderer struct {
	img    *ebiten.Image
	grid   *Grid
	pixels []byte
}

func NewGridRenderer(g *Grid) *GridRenderer {
	w, h := g.Size()

	return &GridRenderer{
		img:    ebiten.NewImage(w, h),
		grid:   g,
		pixels: make([]byte, w*h*4),
	}
}

func (g *GridRenderer) Draw(dst *ebiten.Image) {
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

	g.img.WritePixels(g.pixels)

	dst.DrawImage(g.img, nil)
}

func (g *GridRenderer) Image() *ebiten.Image {
	return g.img
}
