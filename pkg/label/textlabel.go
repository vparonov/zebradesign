package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type TextLabel struct {
	Cell
	Lines int
	Size  float64
	Text  string
	Font  string
}

func NewTextLabel() *TextLabel {
	return &TextLabel{
		Lines: 1,
		Size:  10,
		Text:  "",
		Font:  "",
	}
}

func (c *TextLabel) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("label")
	return b
}
