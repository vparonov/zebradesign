package label

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMMtoPoint(t *testing.T) {
	p := &PageSettings{DPI: 203, Direction: 270, Width: 150, Height: 100}

	points := p.mmToPoint(10)

	assert.Equal(t, 79, points)

	xp, yp := p.toPageCoordinates(0, 0)

	assert.Equal(t, 0, xp)
	assert.Equal(t, p.mmToPoint(p.Width), yp)

	p = &PageSettings{DPI: 203, Direction: 0, Width: 150, Height: 100}

	xp, yp = p.toPageCoordinates(0, 0)

	assert.Equal(t, 0, xp)
	assert.Equal(t, 0, yp)

}
