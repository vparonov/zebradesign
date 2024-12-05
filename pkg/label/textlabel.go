package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type TextLabel struct {
	Cell
	Lines int
	Size  float64
	Text  string
	Font  string
}

func (c *TextLabel) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("label")
	return b
}
