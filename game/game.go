package game

import (
	"dojo-game-of-life/debug"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
	Update()
	Draw(screen *ebiten.Image)
}

type WithLayout interface {
	Layout(int, int) (int, int)
}

type DefaultGame struct {
	game Game
	w, h int
}

func NewDefaultGame(game Game, w, h int) *DefaultGame {
	return &DefaultGame{
		game: game,
		w:    w,
		h:    h,
	}
}

func (g *DefaultGame) Update() error {
	g.game.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	return nil
}

func (g *DefaultGame) Draw(screen *ebiten.Image) {
	g.game.Draw(screen)
	debug.DrawFTPS(screen)
	debug.DrawPrintf(screen, debug.TopRight, "[ALT]+[ENTER] Fullscreen\n[ESC] Exit")
}

func (g *DefaultGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g, ok := g.game.(WithLayout); ok {
		return g.Layout(outsideWidth, outsideHeight)
	}
	return g.w, g.h
}
