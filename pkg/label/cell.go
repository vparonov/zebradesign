package label

import "github.com/vparonov/zebradesign/pkg/zpl"

type CellInterface interface {
	ToZPL(p *PageSettings, b *zpl.ZplBuilder) *zpl.ZplBuilder
}

type Cell struct {
	Type string
	ID   string
	X    float64
	Y    float64
}
