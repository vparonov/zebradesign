package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BarcodeCell struct {
	Cell
	Text                    string
	BarcodeType             string
	ModuleWidth             int     // points
	WToNRatio               float64 // default = 3.0
	Height                  float64 // mm
	Direction               string  // 'N', 'R', 'I', 'B'
	InterpretationLine      bool
	InterpretationLineAbove bool
}

func NewBarcodeCell() *BarcodeCell {
	return &BarcodeCell{
		BarcodeType:             "Code128",
		Text:                    "", //if not provided, the ID will be used as template field (i.e. {{ <<id>> }})
		ModuleWidth:             2,
		WToNRatio:               3.0,
		Direction:               "", // 'N', 'R', 'I', 'B'. '' == default direction for the page settings
		InterpretationLine:      false,
		InterpretationLineAbove: false,
		Height:                  10, // mm, adjust as needed for different barcode heights and widths (in mm)
	}
}

func (c *BarcodeCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	if c.BarcodeType != "Code128" {
		panic("Unsupported barcode type")
	}
	height := p.mmToPoints(c.Height)
	xpage, ypage := p.toPageCoordinates(c.X, c.Y)
	var direction rune
	if c.Direction == "" {
		direction = directionToZPL(p.Direction)
	} else {
		direction = rune(c.Direction[0])
	}

	if direction == 'B' {
		xpage += height
	}

	var value string

	if len(c.Text) > 0 {
		value = c.Text
	} else {
		value = toTemplate(c.ID)
	}
	// ^BY3,3,118^FT162,1160^BCB,,N,N^FD{{ barcode }}^FS
	b.BarCodeFieldDefault(c.ModuleWidth, c.WToNRatio, height).
		FieldTypeset(xpage, ypage).
		Code128BarCode(direction, height, c.InterpretationLine, c.InterpretationLineAbove, false, 'N').
		FieldData(value).FieldSeparator().NewLine()
	return b
}
