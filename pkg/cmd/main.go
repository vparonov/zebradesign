package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/vparonov/zebradesign/pkg/label"
)

func main() {
	inputFileName := flag.String("i", "", "Input label file name")
	pageWidth := flag.Float64("w", 150.0, "Page width [mm] (default: 150)")
	pageHeight := flag.Float64("h", 100.0, "Page width [mm] (default: 100)")
	pageDPI := flag.Float64("d", 203.0, "DPI (default: 203)")
	pageDirection := flag.Int("o", 270, "Page orientation (default: 270)")
	demo := flag.Bool("demo", false, "Generate demo label (default: false)")

	flag.Parse()

	pageSettings := label.NewPageSettings(*pageWidth, *pageHeight, *pageDPI, *pageDirection)
	// open and process the input file
	if *inputFileName == "" {
		panic("Input file name is required")
	}

	// read the label file
	file, err := os.Open(*inputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	var label label.Label
	err = label.UnmarshalJSON(b)

	if err != nil {
		panic(err)
	}

	fmt.Print(label.RenderToPage(pageSettings, *demo))

}
