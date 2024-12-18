package label

import (
	"encoding/json"
	"errors"

	"github.com/vparonov/zebradesign/pkg/zpl"
)

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
^MMT
`
	EPILOG_ZPL = `^PQ1,0,1,Y
^XZ`
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
			i = NewTextCell()
		// the label is a text cell with text property set
		case "label":
			i = NewTextCell()
		case "barcode":
			i = NewBarcodeCell()
		case "box":
			i = NewBoxCell()
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

func (l *Label) RenderToPage(p *PageSettings, demoMode bool) string {
	zplBuilder := zpl.New()

	zplBuilder.RawCode(PROLOG_ZPL)

	if p.Direction == 90 || p.Direction == 270 {
		zplBuilder.PrintWidth(p.mmToPoints(p.Height)).
			PrintLength(p.mmToPoints(p.Width)).
			NewLine()
	} else {
		zplBuilder.PrintWidth(p.mmToPoints(p.Width)).
			PrintLength(p.mmToPoints(p.Height)).
			NewLine()
	}

	for _, c := range l.Cells {
		c.ToZPL(p, zplBuilder, demoMode)
	}
	zplBuilder.RawCode(EPILOG_ZPL)
	return zplBuilder.String()
}
