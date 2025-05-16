package controls

import "github.com/hajimehoshi/ebiten/v2"

type Move struct {
	// Modifier is the key used to activate the move function
	Modifier ebiten.Key

	// tx, ty current translation
	tx, ty float64

	// lx, ly Last known mouse position
	lx, ly float64

	// mState modifier state
	mState bool

	// mbState mouse button state
	mbState bool

	// scale last known scale
	scale float64
}

func NewMove() *Move {
	return &Move{
		Modifier: ebiten.KeySpace,
		scale:    1,
	}
}

// Value returns the current translation
func (m *Move) Value() (float64, float64) {
	mx, my := ebiten.CursorPosition()
	mxf, myf := float64(mx), float64(my)

	if m.mState && m.mbState {
		m.tx += m.lx - mxf
		m.ty += m.ly - myf
	}

	m.lx, m.ly = mxf, myf

	if ebiten.IsKeyPressed(m.Modifier) {
		if !m.mState {
			m.mState = true
			pushCursor(ebiten.CursorShapeMove)
		}
	} else {
		if m.mState {
			m.mState = false
			popCursor(ebiten.CursorShapeMove)
		}

		return m.tx, m.ty
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !m.mbState {
			m.mbState = true
			m.lx, m.ly = mxf, myf
		}
	} else {
		if m.mbState {
			m.mbState = false
			m.lx, m.ly = 0, 0
		}
	}

	return m.tx, m.ty
}

// ValueWithDeltaScaleAt returns the current translation
// with scale correction based on the delta scale (ds)
// and a focus point (fx, fy) in the screen
func (m *Move) ValueWithDeltaScaleAt(ds float64, fx, fy float64) (float64, float64) {
	if ds == 0 {
		return m.Value()
	}

	m.Value() // update tx, ty

	oldScale := m.scale
	newScale := oldScale + ds

	m.tx += -(m.tx + fx) * (1 - newScale/oldScale)
	m.ty += -(m.ty + fy) * (1 - newScale/oldScale)

	m.scale = newScale
	return m.tx, m.ty
}
