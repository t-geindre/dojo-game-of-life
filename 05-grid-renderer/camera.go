package d05

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Camera struct {
	// curScale current scale (smoothed)
	// tarScale target scale (clamped between minScale and maxScale)
	// smthScale is the scale smoothing factor (0-1)
	curScale, tarScale, smthScale, minScale, maxScale float64

	// src is this image source used by the camera
	src *ebiten.Image
}

func NewCamera() *Camera {
	return &Camera{
		smthScale: .25,
		curScale:  1,
		tarScale:  1,
		minScale:  .5,
		maxScale:  5,
	}
}

func (c *Camera) Update() {
	// Scale smoothing
	ds := c.tarScale - c.curScale
	if math.Abs(ds) < 0.001 {
		c.curScale = c.tarScale
	} else {
		c.curScale += ds * c.smthScale
	}
}

// Draw draws the camera image to the destination image
func (c *Camera) Draw(dst *ebiten.Image) {
	if c.src == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(c.curScale, c.curScale)
	dst.DrawImage(c.src, op)
}

// Attach sets the image source for the camera
func (c *Camera) Attach(src *ebiten.Image) {
	c.src = src
}

// Detach clears the image source for the camera
func (c *Camera) Detach() {
	c.src = nil
}

// Scale sets the target zoom level
// .1 = 10%, 1 = 100%, 2 = 200%, etc
func (c *Camera) Scale(f float64) {
	c.tarScale = math.Min(math.Max(c.tarScale+f, c.minScale), c.maxScale)
}

// CurrentScale returns the current zoom level
func (c *Camera) CurrentScale() float64 {
	return c.curScale
}

// ScaleSmoothing sets the zoom smoothing factor, 0-1 (amount of zoom applied on each update)
func (c *Camera) ScaleSmoothing(f float64) {
	c.smthScale = f
}
