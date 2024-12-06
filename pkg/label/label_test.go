package label

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vparonov/zebradesign/pkg/zpl"
)

func TestMMtoPoint(t *testing.T) {
	p := &PageSettings{DPI: 203, Direction: 270, Width: 150, Height: 100}

	points := p.mmToPoints(100)

	assert.Equal(t, 799, points)

	xp, yp := p.toPageCoordinates(0, 0)

	assert.Equal(t, 0, xp)
	assert.Equal(t, p.mmToPoints(p.Width), yp)

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
	assert.Equal(t, 1, cell.Lines)
	assert.Equal(t, "", cell.Font)

}

func TestDemarchalLabel(t *testing.T) {
	p := &PageSettings{DPI: 203, Direction: 270, Width: 150, Height: 100}

	labelJSON := `{
		"Cells":[
			{
	    		"Type": "text", 
        		"ID": "cell1",
        		"X": 0,"Y": 0,
				"Width": 20, 
				"Text": "Hello, World!",
				"Font": "0","Lines": 1,
				"Size": 5
			},
			{
				"Type": "box", 
				"ID": "cell3",
                "X": 0,"Y": 0,
                "Width": 20,
                "Height": 10
            },{
				"Type": "box", 
				"ID": "cell2",
                "X": 10,"Y": 20,
                "Width": 50,
                "Height": 10
            },
			{
	    		"Type": "barcode", 
        		"ID": "cell1",
        		"X": 10,"Y": 20,
				"BarcodeType": "Code128",
				"Size": 10
			},
			{
	    		"Type": "barcode", 
        		"ID": "cell33",
        		"X": 80,"Y": 20,
				"Direction": "N",
				"BarcodeType": "Code128",
				"Size": 10
			}
		]
}`
	//				"Direction": "N",

	var label Label
	err := label.UnmarshalJSON([]byte(labelJSON))

	assert.Nil(t, err)

	cell := label.Cells[0]
	zplResult := cell.ToZPL(p, zpl.New()).String()
	assert.NotEmpty(t, zplResult)

	cell = label.Cells[1]
	zplResult = cell.ToZPL(p, zpl.New()).String()
	assert.NotEmpty(t, zplResult)

	page := label.RenderToPage(p)

	assert.NotEmpty(t, page)
	fmt.Print(page)
}
