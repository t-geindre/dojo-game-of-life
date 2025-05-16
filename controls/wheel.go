package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// Wheel is a struct that handles mouse wheel
type Wheel struct {
	// Min & Max are the min and max values
	// .1 is 10% zoom 10 is 10 times bigger
	Min, Max float64

	// Rate is the wheel sensitivity
	// 0-1, 0 to 100%
	Rate float64

	// Speed is the smoothing speed, time spent to reach the target
	// 0-1, 0 to 100%
	Speed float64

	// Modifier is the key used to activate the mouse wheel control
	Modifier ebiten.Key

	// Cursor is the cursor shape used when the modifier is pressed
	Cursor ebiten.CursorShapeType

	// Modifiers states
	mState bool

	// target zoom level
	target float64

	// current current value
	current float64

	// delta last delta value
	delta float64

	// atx, aty zoom geo target
	atx, aty float64
}

func NewWheel() *Wheel {
	return &Wheel{
		target:   1,
		current:  1,
		Min:      .5,
		Max:      20,
		Rate:     0.3,
		Speed:    .2,
		Modifier: ebiten.KeyControl,
		Cursor:   ebiten.CursorShapeNSResize,
	}
}

// Value current value, smoothed
func (w *Wheel) Value() float64 {
	if ebiten.IsKeyPressed(w.Modifier) {
		if !w.mState {
			w.mState = true
			pushCursor(w.Cursor)
		}
		_, mwy := ebiten.Wheel()

		if mwy != 0 {
			mx, my := ebiten.CursorPosition()
			w.atx, w.aty = float64(mx), float64(my)
		}

		w.target = math.Min(w.Max, math.Max(w.target+mwy*w.Rate, w.Min))
	} else if w.mState {
		w.mState = false
		popCursor(w.Cursor)
	}

	d := w.target - w.current

	if math.Abs(d) < 0.001 {
		w.current = w.target
		w.delta = d
	} else {
		w.delta = d * w.Speed
		w.current += w.delta
	}

	return w.current
}

// SetValue sets the zoom value
func (w *Wheel) SetValue(v float64) {
	w.target = math.Min(w.Max, math.Max(v, w.Min))
	w.current = w.target
}

// Delta returns the last delta value
func (w *Wheel) Delta() float64 {
	return w.delta
}

// Focus gets the zoom coordinates
func (w *Wheel) Focus() (float64, float64) {
	return w.atx, w.aty
}
