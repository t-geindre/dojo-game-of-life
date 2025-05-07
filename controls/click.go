package controls

type Draw struct {
}

func NewDraw() *Draw {
	return &Draw{}
}

func (c *Draw) update() {

}

func (c *Draw) Clicked() bool {
	return false
}

func (c *Draw) Where(s, tx, ty float64) (float64, float64) {
	return 0, 0
}
