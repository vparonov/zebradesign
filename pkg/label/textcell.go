package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type TextCell struct {
	Cell
	Lines int
	Size  float64
	Font  string
}

func NewTextCell() *TextCell {
	return &TextCell{
		Lines: 1,
		Size:  10,
		Font:  "",
	}
}

func (c *TextCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("text")
	return b
}
