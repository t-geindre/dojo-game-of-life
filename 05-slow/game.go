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
	zoom   *controls.Wheel2
	scale  float64
	camera *Camera
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)
	grid.Randomize(.3)

	buff := ebiten.NewImage(w, h)

	cam := NewCamera()
	cam.Attach(buff)

	zoom := controls.NewWheel2(ebiten.KeyControl, ebiten.CursorShapeNSResize)
	zoom.Attach(func(w *controls.Wheel2) {
		cam.Scale(w.Value())
	})

	return &Game{
		pixels: make([]byte, w*h*4),
		grid:   grid,
		buff:   buff,
		zoom:   zoom,
		scale:  1,
		camera: cam,
	}
}

func (g *Game) Update() {
	g.zoom.Update()
	g.camera.Update()

	g.grid.NextState()
	g.UpdatePixels()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.buff.WritePixels(g.pixels)

	g.camera.Draw(screen)

	debug.DrawPrintf(
		screen, debug.BottomRight,
		"[CTL]+[MW] Zoom (%0.f%%)",
		g.camera.CurrentScale()*100,
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
