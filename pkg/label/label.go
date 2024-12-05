package label

import (
	"encoding/json"
	"errors"
)

type PageSettings struct {
	Width     float64
	Height    float64
	DPI       float64
	Direction int
}

type CellInterface interface {
	ToZPL() string
}

type Cell struct {
	Type string
	ID   string
	X    float64
	Y    float64
	Size float64
	BL   bool
	BR   bool
	BT   bool
	BB   bool
}

type TextCell struct {
	Cell
	Lines int
	Text  string
	Font  string
}

type BarcodeCell struct {
	Cell
	BarcodeType string
}

type Label struct {
	Cells    []CellInterface   `json:"-"`
	RawCells []json.RawMessage `json:"Cells"`
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

func (l *Label) UnmarshalJSON(data []byte) error {
	type label Label

	err := json.Unmarshal(data, (*label)(l))
	if err != nil {
		return err
	}

	for _, raw := range l.RawCells {
		var c Cell
		err = json.Unmarshal(raw, &c)
		if err != nil {
			return err
		}
		var i CellInterface
		switch c.Type {
		case "text":
			i = &TextCell{}
		case "barcode":
			i = &BarcodeCell{}
		default:
			return errors.New("unknown cell type")
		}
		err = json.Unmarshal(raw, i)
		if err != nil {
			return err
		}
		l.Cells = append(l.Cells, i)
	}
	return nil
}

func (l *Label) RenderToPage(pageSettings *PageSettings) string {

	return ""
}

func (c *TextCell) ToZPL() string {
	return "text"
}

func (c *BarcodeCell) ToZPL() string {
	return "barcode"
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
