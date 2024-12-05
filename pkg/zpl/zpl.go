package zpl

import (
	"fmt"
	"strings"
)

type ZplBuilder struct {
	sb strings.Builder
}

func New() *ZplBuilder {
	return &ZplBuilder{}
}

func (b *ZplBuilder) RawCode(rawCode string) *ZplBuilder {
	b.sb.WriteString(rawCode)
	return b
}
func (b *ZplBuilder) PrintLength(points int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^LL%d", points))
	return b
}
func (b *ZplBuilder) PrintWidth(points int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^PW%d", points))
	return b
}

func (b *ZplBuilder) FieldOrigin(x, y int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^FO%d,%d", x, y))
	return b
}

func (b *ZplBuilder) GraphicBox(width, height, borderWidth int, lineColor bool, rounding int) *ZplBuilder {
	c := 'W'
	if lineColor {
		c = 'B'
	}
	b.sb.WriteString(fmt.Sprintf("^GB%d,%d,%d,%c,%d", width, height, borderWidth, c, rounding))
	return b
}

func (b *ZplBuilder) FieldSeparator() *ZplBuilder {
	b.sb.WriteString("^FS")
	return b
}

func (b *ZplBuilder) NewLine() *ZplBuilder {
	b.sb.WriteString("\r\n")
	return b
}
func (b *ZplBuilder) String() string {
	return b.sb.String()
}
