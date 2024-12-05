package label

type TextCell struct {
	Cell
	Lines int
	Text  string
	Font  string
}

func (c *TextCell) ToZPL(p *PageSettings) string {
	return "text"
}
