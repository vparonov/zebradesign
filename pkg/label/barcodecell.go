package label

type BarcodeCell struct {
	Cell
	BarcodeType string
}

func (c *BarcodeCell) ToZPL(p *PageSettings) string {
	return "barcode"
}
