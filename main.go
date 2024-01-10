package main

import "fmt"

type ColorValue struct {
	Foreground int
	Background int
}

type AvailableColors int

const (
	Black AvailableColors = iota + 1
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

type AvailableStyles int

const (
	Bold AvailableStyles = iota + 1
	Italic
	Underline
)

type StyleCodes map[AvailableStyles]int

var styles StyleCodes = StyleCodes{
	Bold:      1,
	Italic:    3,
	Underline: 4,
}

type ColorsCodes map[AvailableColors]ColorValue

var colors ColorsCodes = ColorsCodes{
	Black:   ColorValue{30, 40},
	Red:     ColorValue{31, 41},
	Green:   ColorValue{32, 42},
	Yellow:  ColorValue{33, 43},
	Blue:    ColorValue{34, 44},
	Magenta: ColorValue{35, 45},
	Cyan:    ColorValue{36, 46},
	White:   ColorValue{37, 47},

	BrightBlack:   ColorValue{90, 100},
	BrightRed:     ColorValue{91, 101},
	BrightGreen:   ColorValue{92, 102},
	BrightYellow:  ColorValue{93, 103},
	BrightBlue:    ColorValue{94, 104},
	BrightMagenta: ColorValue{95, 105},
	BrightCyan:    ColorValue{96, 106},
	BrightWhite:   ColorValue{97, 107},
}

const escape = "\033"
const reset = escape + "[0m"

type Colorizer struct{}

func (c *Colorizer) MountString(str string, code int) string {
	return fmt.Sprintf("%s[%dm%s%s", escape, code, str, reset)
}
func (c *Colorizer) Black(i string) string {
	return c.MountString(i, colors[Black].Foreground)
}

func (c *Colorizer) Red(i string) string {
	return c.MountString(i, colors[Red].Foreground)
}
func (c *Colorizer) Green(i string) string {
	return c.MountString(i, colors[Green].Foreground)
}
func (c *Colorizer) Yellow(i string) string {
	return c.MountString(i, colors[Yellow].Foreground)
}
func (c *Colorizer) Blue(i string) string {
	return c.MountString(i, colors[Blue].Foreground)
}
func (c *Colorizer) Magenta(i string) string {
	return c.MountString(i, colors[Magenta].Foreground)
}
func (c *Colorizer) Cyan(i string) string {
	return c.MountString(i, colors[Cyan].Foreground)
}
func (c *Colorizer) White(i string) string {
	return c.MountString(i, colors[White].Foreground)
}
func (c *Colorizer) BgBlack(i string) string {
	return c.MountString(i, colors[Black].Background)
}
func (c *Colorizer) BgRed(i string) string {
	return c.MountString(i, colors[Red].Background)
}
func (c *Colorizer) BgGreen(i string) string {
	return c.MountString(i, colors[Green].Background)
}
func (c *Colorizer) BgYellow(i string) string {
	return c.MountString(i, colors[Yellow].Background)
}
func (c *Colorizer) BgBlue(i string) string {
	return c.MountString(i, colors[Blue].Background)
}
func (c *Colorizer) BgMagenta(i string) string {
	return c.MountString(i, colors[Magenta].Background)
}
func (c *Colorizer) BgCyan(i string) string {
	return c.MountString(i, colors[Cyan].Background)
}
func (c *Colorizer) BgWhite(i string) string {
	return c.MountString(i, colors[White].Background)
}

func (c *Colorizer) Bold(i string) string {
	return c.MountString(i, styles[Bold])
}

func (c *Colorizer) Italic(i string) string {
	return c.MountString(i, styles[Italic])
}

func (c *Colorizer) Underline(i string) string {
	return c.MountString(i, styles[Underline])
}

func (c *Colorizer) includeStyle(str string, code int) string {
	return fmt.Sprintf("%s[%dm%s", escape, code, str)
}

type Options struct {
	Color AvailableColors
	Bg    AvailableColors
	Style AvailableStyles
}

func (c *Colorizer) Compose(str string, options Options) string {
	out := str + reset

	if options.Bg != 0 {
		out = c.includeStyle(out, colors[options.Bg].Background)
	}

	if options.Color != 0 {
		out = c.includeStyle(out, colors[options.Color].Foreground)
	}

	if options.Style != 0 {
		out = c.includeStyle(out, styles[options.Style])
	}

	return out
}

func main() {
	var c Colorizer
	fmt.Println("Colors:", c.Red("red"), c.Green("green"), c.Blue("blue"), c.BgWhite(c.Black("black")), "normal")
	fmt.Println("Styles:", c.Bold("bold"), c.Underline("underline"), c.Italic("italic"), "normal")
	fmt.Println("Composed:", c.Compose("composed", Options{Color: Black, Bg: BrightWhite, Style: Bold}))

}
