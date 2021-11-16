package helpers

type Dimension struct {
	W, H int
}

func NewDimensions(w, h int) Dimension {
	return Dimension{w, h}
}

func (d *Dimension) Set(w, h int) {
	d.W = w
	d.H = h
}
