package controls

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type Zoom struct {
	target  float64
	current float64
	last    time.Time
}

func NewZoom() *Zoom {
	return &Zoom{
		target:  1,
		current: 1,
		last:    time.Now(),
	}
}

func (z *Zoom) Zoom() float64 {
	_, mwy := ebiten.Wheel()
	if mwy != 0 {
		z.last = time.Now()
	}
	f := float64(time.Since(z.last).Milliseconds())
	z.target += float64(mwy) / math.Max(1, f) / 50
	z.current += (z.target - z.current) / 5

	return z.target - z.current
}
