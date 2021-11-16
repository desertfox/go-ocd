package helpers

type Dimension struct {
	W, H int
}

func NewDimensions() Dimension {
	return Dimension{int(0), int(0)}
}

func (d *Dimension) Set(w, h int) {
	d.W = w
	d.H = h
}
