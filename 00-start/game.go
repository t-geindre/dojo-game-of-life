package d00

import (
	"dojo-game-of-life/game"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

//go:embed shader.kage
var shaderSrc []byte

type Game struct {
	shader *ebiten.Shader
	start  time.Time
}

func NewGame(_, _ int) game.Game {
	shader, err := ebiten.NewShader(shaderSrc)
	if err != nil {
		panic(err)
	}

	return &Game{
		shader: shader,
		start:  time.Now(),
	}
}

func (g *Game) Update() {
}

func (g *Game) Draw(screen *ebiten.Image) {
	bd := screen.Bounds()
	screen.DrawRectShader(bd.Dx(), bd.Dy(), g.shader, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]any{
			"Time": float32(time.Since(g.start).Seconds()),
		},
	})
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
