package controls

import "github.com/hajimehoshi/ebiten/v2"

type Move struct {
	lx, ly int
	md     bool
}

func NewMove() *Move {
	return &Move{}
}

func (m *Move) Move() (float64, float64) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		if !m.md {
			m.lx, m.ly = x, y
			m.md = true
			ebiten.SetCursorShape(ebiten.CursorShapeMove)
		} else {
			mx, my := x-m.lx, y-m.ly
			m.lx, m.ly = x, y
			return float64(mx), float64(my)
		}
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		m.md = false
	}

	x, y := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		x = 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		x = -10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		y = 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		y = -10
	}

	return x, y
}
