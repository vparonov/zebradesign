package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BoxCell struct {
	Cell
	Width       float64
	Height      float64
	BorderWidth float64
	BL          bool
	BR          bool
	BT          bool
	BB          bool
}

func NewBoxCell() *BoxCell {
	return &BoxCell{
		Width:       10,
		Height:      10,
		BorderWidth: 0.250,
		BL:          true,
		BR:          true,
		BT:          true,
		BB:          true,
	}
}

func (c *BoxCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	if !(c.BL && c.BR && c.BT && c.BB) {
		panic("Box cell must have all sides set to true")
	}

	var xpage, ypage, wpage, hpage int
	xpage, ypage = p.toPageCoordinates(c.X, c.Y)
	wpage = p.mmToPoints(c.Width)
	hpage = p.mmToPoints(c.Height)
	if p.Direction == 270 {
		ypage -= wpage
		wpage, hpage = hpage, wpage
		//xpage -= wpage
	}
	//^FO31,1173^GB678,0,2^FS
	b.FieldOrigin(xpage, ypage).
		GraphicBox(wpage, hpage, p.mmToPoints(c.BorderWidth), true, 0).
		FieldSeparator().NewLine()

	return b
}
