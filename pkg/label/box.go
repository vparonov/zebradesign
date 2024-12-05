package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BoxCell struct {
	Cell
}

func (c *BoxCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("box")
	return b
}
