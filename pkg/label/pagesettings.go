package label

import "math"

type PageSettings struct {
	Width     float64
	Height    float64
	DPI       float64
	Direction int
}

func (p *PageSettings) mmToPoints(mm float64) int {
	return int(math.Round(p.DPI * mm / 25.4))
}

func (p *PageSettings) toPageCoordinates(x, y float64) (xPage int, yPage int) {
	var xPageMM float64
	var yPageMM float64

	switch p.Direction {
	case 0:
		xPageMM = x
		yPageMM = y
	case 90:
		xPageMM = y
		yPageMM = x
	case 180:
		xPageMM = p.Width - x
		yPageMM = p.Height - y
	case 270:
		xPageMM = y
		yPageMM = p.Width - x
	default:
		panic("Invalid direction")
	}

	xPage = p.mmToPoints(xPageMM)
	yPage = p.mmToPoints(yPageMM)

	return
}
