package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BoxCell struct {
	Cell
	BL bool
	BR bool
	BT bool
	BB bool
}

func NewBoxCell() *BoxCell {
	return &BoxCell{
		BL: true,
		BR: true,
		BT: true,
		BB: true,
	}
}

func (c *BoxCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("box")
	return b
}
