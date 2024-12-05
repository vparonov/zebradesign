package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BarcodeCell struct {
	Cell
	ModuleWidth float64 // mm
	WToNRatio   float64 // default = 3.0
	Height      float64 // mm
	BarcodeType string
}

func (c *BarcodeCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	b.RawCode("barcode")
	return b
}
