package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type TextCell struct {
	Cell
	Text              string
	Lines             int
	Size              float64
	Width             float64
	Font              string
	TextJustification rune
}

func NewTextCell() *TextCell {
	return &TextCell{
		Lines:             1,
		Size:              10,
		Width:             100,
		TextJustification: 'L',
		Font:              "",
		Text:              "",
	}
}

func (c *TextCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	//^FT65,400^A@B,45,45,{{ font_name }}^FH\^CI17^F8^FD{{ pack }}^FS^CI0
	totalHeight := c.Lines * p.mmToPoints(c.Size)

	xpage, ypage := p.toPageCoordinates(c.X, c.Y)

	if p.Direction == 270 {
		xpage += totalHeight
	}

	b.FieldTypeset(xpage, ypage)

	// Font could be A-Z, 0-9
	if len(c.Font) == 1 {
		b.FontByName(rune(c.Font[0]), directionToZPL(p.Direction), p.mmToPoints(c.Size), p.mmToPoints(c.Size))
	} else {
		// or file name from the printer's font directory
		b.FontByFileName(directionToZPL(p.Direction), p.mmToPoints(c.Size), p.mmToPoints(c.Size), c.Font)
	}

	var value string
	if len(c.Text) > 0 {
		value = c.Text
	} else {
		value = toTemplate(c.ID)
	}

	b.FieldBlock(p.mmToPoints(c.Width), c.Lines, 0, c.TextJustification, 0).
		FieldHexadecimalIndicator('\\').
		CyrCharset().
		FieldData(value).
		FieldSeparator().
		ResetCharset().
		NewLine()
	return b
}
