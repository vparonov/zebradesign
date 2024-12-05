package label

type PageSettings struct {
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	DPI       float64 `json:"dpi"`
	Direction int     `json:"direction"`
}

type Cell struct {
	ID    string  `json:"id"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Lines int     `json:"lines"`
	Text  string  `json:"text"`
	Size  float64 `json:"size"`
	Font  string  `json:"font"`
	BL    bool    `json:"bl"`
	BR    bool    `json:"br"`
	BT    bool    `json:"bt"`
	BB    bool    `json:"bb"`
}

type Label struct {
	Cells []Cell
}

const (
	PROLOG_ZPL = `CT~~CD,~CC^~CT~
^XA
~TA000~JSN^LT0^MNW^MTT^PON^PMN^LH0,0^JMA^PR8,8~SD15^JUS^LRN^CI0
^XZ
^XA
`
	EPILOG_ZPL = `^PQ1,0,1,Y
^XZ'`
)

func (l *Label) RenderToTemplate() string {
	return ""
}

func (p *PageSettings) mmToPoint(mm float64) int {
	return int(p.DPI * mm / 25.4)
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

	xPage = p.mmToPoint(xPageMM)
	yPage = p.mmToPoint(yPageMM)

	return
}
