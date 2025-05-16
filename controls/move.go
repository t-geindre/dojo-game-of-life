package controls

import "github.com/hajimehoshi/ebiten/v2"

type Move struct {
	// Modifier is the key used to activate the move function
	Modifier ebiten.Key

	// mState modifier state
	mState bool

	// tx, ty current translation
	tx, ty float64

	// lx, ly Last known mouse position
	lx, ly float64
}

func NewMove() *Move {
	return &Move{
		Modifier: ebiten.KeySpace,
	}
}

func (m *Move) Value() (float64, float64) {
	if ebiten.IsKeyPressed(m.Modifier) {
		mx, my := ebiten.CursorPosition()
		mxf, myf := float64(mx), float64(my)

		if !m.mState {
			m.mState = true
			pushCursor(ebiten.CursorShapeMove)
			m.lx, m.ly = mxf, myf
			return m.tx, m.ty
		}

		m.tx += m.lx - mxf
		m.ty += m.ly - myf

		m.lx, m.ly = mxf, myf

	} else if m.mState {
		m.mState = false
		popCursor(ebiten.CursorShapeMove)
	}

	return m.tx, m.ty
}
