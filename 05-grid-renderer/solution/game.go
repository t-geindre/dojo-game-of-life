package d05s

import (
	"dojo-game-of-life/controls"
	"dojo-game-of-life/debug"
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid     *Grid
	renderer *GridRenderer
	zoom     *controls.Wheel2
	camera   *Camera
}

func NewGame(w, h int) game.Game {
	grid := NewGrid(w, h)
	grid.Randomize(.3)

	renderer := NewGridRenderer(grid)

	cam := NewCamera()
	cam.Attach(renderer.Image())

	zoom := controls.NewWheel2(ebiten.KeyControl, ebiten.CursorShapeNSResize)
	zoom.Attach(func(w *controls.Wheel2) {
		cam.Scale(w.Value())
	})

	return &Game{
		grid:     grid,
		zoom:     zoom,
		renderer: renderer,
		camera:   cam,
	}
}

func (g *Game) Update() {
	g.zoom.Update()
	g.camera.Update()
	g.grid.NextState()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen)
	g.camera.Draw(screen)

	debug.DrawPrintf(
		screen, debug.BottomRight,
		"[CTL]+[MW] Zoom (%0.f%%)",
		g.camera.CurrentScale()*100,
	)
}
