package label

type CellInterface interface {
	ToZPL(p *PageSettings) string
}

type Cell struct {
	Type string
	ID   string
	X    float64
	Y    float64
	Size float64
	BL   bool
	BR   bool
	BT   bool
	BB   bool
}
