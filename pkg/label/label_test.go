package label

import (
	"encoding/json"
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

func TestDemarshalCells(t *testing.T) {
	singleTextCell := `
{
	"Type": "text", 
	"ID": "cell1",
	"X": 10,
    "Y": 20,
    "Size": 10,
    "BL": true,
    "BR": true,
    "BT": true,
    "BB": true,
	"Lines": 1,
	"Text": "Hello, World!",
    "Font": ""
}
`

	var cell TextCell

	err := json.Unmarshal([]byte(singleTextCell), &cell)
	assert.Nil(t, err)

	assert.Equal(t, "cell1", cell.ID)
	assert.Equal(t, float64(10), cell.X)
	assert.Equal(t, float64(20), cell.Y)
	assert.Equal(t, float64(10), cell.Size)
	assert.True(t, cell.BL)
	assert.True(t, cell.BR)
	assert.True(t, cell.BT)
	assert.True(t, cell.BB)
	assert.Equal(t, 1, cell.Lines)
	assert.Equal(t, "Hello, World!", cell.Text)
	assert.Equal(t, "", cell.Font)

}

func TestDemarchalLabel(t *testing.T) {
	labelJSON := `{"Cells":[{
	    "Type": "text", 
        "ID": "cell1",
        "X": 10,"Y": 20,"Text": "Hello, World!","Font": "","Lines": 1,"Size": 10}]}
`

	var label Label
	err := label.UnmarshalJSON([]byte(labelJSON))

	assert.Nil(t, err)

	for _, cell := range label.Cells {
		zpl := cell.ToZPL()
		assert.Equal(t, "text", zpl)
	}
}
