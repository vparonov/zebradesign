package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type BarcodeCell struct {
	Cell
	BarcodeType string
	ModuleWidth float64 // mm
	WToNRatio   float64 // default = 3.0
	Height      float64 // mm
}

func NewBarcodeCell() *BarcodeCell {
	return &BarcodeCell{
		BarcodeType: "Code128",
		ModuleWidth: 0.375,
		WToNRatio:   3.0,
		Height:      10, // mm, adjust as needed for different barcode heights and widths (in mm)
	}
}
func (c *BarcodeCell) ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder {
	if c.BarcodeType != "Code128" {
		panic("Unsupported barcode type")
	}

	// ^BY3,3,118^FT162,1160^BCB,,N,N^FD{{ barcode }}^FS
	b.BarCodeFieldDefault(p.mmToPoint(c.ModuleWidth), c.WToNRatio, p.mmToPoint(c.Height))
	return b
}
