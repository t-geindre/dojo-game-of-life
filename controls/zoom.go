package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	ZoomMin = 0.5
	ZoomMax = 20
)

type Zoom struct {
	target  float64
	current float64
}

func NewZoom() *Zoom {
	return &Zoom{
		target:  1,
		current: 1,
	}
}

func (z *Zoom) Zoom() float64 {
	_, mwy := ebiten.Wheel()
	if mwy != 0 {
		t := z.target + float64(mwy)
		if t < 1 {
			t = z.target + float64(mwy)*0.5
		}
		z.target = t
	}

	z.target = math.Min(20, math.Max(z.target, -.5))
	z.current += (z.target - z.current) / 10

	return 1 + z.current
}
