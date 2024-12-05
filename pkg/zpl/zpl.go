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

func (b *ZplBuilder) FieldTypeset(x, y int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^FT%d,%d", x, y))
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

func (b *ZplBuilder) BarCodeFieldDefault(w int, r float64, h int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^BY%d,%1.1f,%d", w, r, h))
	return b
}

func (b *ZplBuilder) Code128BarCode(o rune, h int, f bool, g bool, e bool, m rune) *ZplBuilder {
	f_v := 'N'
	g_v := 'N'
	e_v := 'N'
	if f {
		f_v = 'Y'
	}
	if g {
		g_v = 'Y'
	}
	if e {
		e_v = 'Y'
	}
	b.sb.WriteString(fmt.Sprintf("^BC%c,%d,%c,%c,%c,%c", o, h, f_v, g_v, e_v, m))
	return b
}

func (b *ZplBuilder) FontByFileName(o rune, h int, w int, fontFileName string) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^A@%c,%d,%d,%s", o, h, w, fontFileName))
	return b
}

func (b *ZplBuilder) FontByName(f rune, o rune, h int, w int) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^A%c%c,%d,%d", f, o, h, w))
	return b
}

func (b *ZplBuilder) FieldHexadecimalIndicator(c rune) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^FH%c", c))
	return b
}

func (b *ZplBuilder) FieldData(data string) *ZplBuilder {
	b.sb.WriteString(fmt.Sprintf("^FD%s", data))
	return b
}

func (b *ZplBuilder) FieldSeparator() *ZplBuilder {
	b.sb.WriteString("^FS")
	return b
}

func (b *ZplBuilder) CyrCharset() *ZplBuilder {
	b.sb.WriteString("^CI17^F8") // magic ...... sorry
	return b
}

// CI0
func (b *ZplBuilder) ResetCharset() *ZplBuilder {
	b.sb.WriteString("^CI0") // magic ...... sorry
	return b
}

func (b *ZplBuilder) NewLine() *ZplBuilder {
	b.sb.WriteString("\r\n")
	return b
}
func (b *ZplBuilder) String() string {
	return b.sb.String()
}
