package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Wheel2 struct {
	// rate wheel sensitivity
	rate float64

	// value current value
	value float64

	// modifier key to activate the wheel
	modifier ebiten.Key

	// mState modifier state (pressed or not)
	mState bool

	// modifier cursor shape when modifier is pressed
	cursor ebiten.CursorShapeType

	// attached function to call when the wheel is used
	attached func(*Wheel2)

	// fx & fy position of the last wheel use
	fx, fy float64
}

func NewWheel2(modifier ebiten.Key, cursor ebiten.CursorShapeType) *Wheel2 {
	return &Wheel2{
		rate:     .1,
		modifier: modifier,
		cursor:   cursor,
	}
}

// Update polls the mouse wheel and modifier key state.
// If the modifier key is pressed, the wheel value is updated and
// cursor shape is set to the modifier cursor shape accordingly.
func (w *Wheel2) Update() {
	w.value = 0

	if ebiten.IsKeyPressed(w.modifier) {
		if !w.mState {
			pushCursor(w.cursor)
			w.mState = true
		}
		w.pullValue()
		return
	}

	if w.mState {
		popCursor(w.cursor)
		w.mState = false
	}
}

// Value returns the current wheel value since the last update
func (w *Wheel2) Value() float64 {
	return w.value
}

// Focus return the last position where the wheel was used
func (w *Wheel2) Focus() (float64, float64) {
	return w.fx, w.fy
}

// Attach sets the function to call when the wheel is used
func (w *Wheel2) Attach(f func(*Wheel2)) {
	w.attached = f
}

// Detach clears the function to call when the wheel is used
func (w *Wheel2) Detach() {
	w.attached = nil
}

// SetRate sets the wheel sensitivity
func (w *Wheel2) SetRate(r float64) {
	w.rate = r
}

func (w *Wheel2) pullValue() {
	_, my := ebiten.Wheel()
	w.value = my * w.rate

	if w.value != 0 {
		mpy, mpx := ebiten.CursorPosition()
		w.fx, w.fy = float64(mpy), float64(mpx)
	}

	if w.attached != nil {
		w.attached(w)
	}
}
