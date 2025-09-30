package styles

import (
	"fmt"
	"image/color"
)

type percentage struct{ value float64 }

func (p *percentage) Value() string { return fmt.Sprintf("%.2f%%", p.value) }

func Percentage(value float64) *percentage { return &percentage{value: value} }

type absoluteSize struct{ value string }

func (a *absoluteSize) Value() string { return a.value }

func XXSmall() *absoluteSize  { return &absoluteSize{value: "xx-small"} }
func XSmall() *absoluteSize   { return &absoluteSize{value: "x-small"} }
func Small() *absoluteSize    { return &absoluteSize{value: "small"} }
func Medium() *absoluteSize   { return &absoluteSize{value: "medium"} }
func Large() *absoluteSize    { return &absoluteSize{value: "large"} }
func XLarge() *absoluteSize   { return &absoluteSize{value: "x-large"} }
func XXLarge() *absoluteSize  { return &absoluteSize{value: "xx-large"} }
func XXXLarge() *absoluteSize { return &absoluteSize{value: "xxx-large"} }

type angle struct {
	value float32
	unit  string
}

func (a *angle) Value() string { return fmt.Sprintf("%.2f%s", a.value, a.unit) }

func Degree(degree float32) *angle   { return &angle{value: degree, unit: "deg"} }
func Gradian(gradian float32) *angle { return &angle{value: gradian, unit: "grad"} }
func Radian(radian float32) *angle   { return &angle{value: radian, unit: "rad"} }
func Turn(turn float32) *angle       { return &angle{value: turn, unit: "turn"} }

type baselinePosition struct {
	first bool
	last  bool
}

func (b *baselinePosition) Value() string {
	if b.first {
		return "first baseline"
	}

	if b.last {
		return "last baseline"
	}

	return "baseline"
}

func Baseline() *baselinePosition      { return &baselinePosition{} }
func FirstBaseline() *baselinePosition { return &baselinePosition{first: true} }
func LastBaseline() *baselinePosition  { return &baselinePosition{last: true} }

type blendMode struct {
	value string
}

func (b *blendMode) Value() string {
	return b.value
}

func NormalBlend() *blendMode     { return &blendMode{value: "normal"} }
func MultiplyBlend() *blendMode   { return &blendMode{value: "multiply"} }
func ScreenBlend() *blendMode     { return &blendMode{value: "screen"} }
func OverlayBlend() *blendMode    { return &blendMode{value: "overlay"} }
func DarkenBlend() *blendMode     { return &blendMode{value: "darken"} }
func LightenBlend() *blendMode    { return &blendMode{value: "lighten"} }
func ColorDodgeBlend() *blendMode { return &blendMode{value: "color-dodge"} }
func ColorBurnBlend() *blendMode  { return &blendMode{value: "color-burn"} }
func HardLightBlend() *blendMode  { return &blendMode{value: "hard-light"} }
func SoftLightBlend() *blendMode  { return &blendMode{value: "soft-light"} }
func DifferenceBlend() *blendMode { return &blendMode{value: "difference"} }
func ExclusionBlend() *blendMode  { return &blendMode{value: "exclusion"} }
func HueBlend() *blendMode        { return &blendMode{value: "hue"} }
func SaturationBlend() *blendMode { return &blendMode{value: "saturation"} }
func ColorBlend() *blendMode      { return &blendMode{value: "color"} }
func LuminosityBlend() *blendMode { return &blendMode{value: "luminosity"} }

type calc struct {
	value string
}

func (c *calc) Value() string {
	return c.value
}

func E() *calc           { return &calc{value: "e"} }
func Pi() *calc          { return &calc{value: "pi"} }
func Infinity() *calc    { return &calc{value: "infinity"} }
func NegInfinity() *calc { return &calc{value: "-infinity"} }
func NaN() *calc         { return &calc{value: "NaN"} }

type ratio struct {
	first  uint
	second uint
}

func (r *ratio) Value() string {
	return fmt.Sprintf("%d / %d", r.first, r.second)
}

func Ratio(first uint, second uint) *ratio {
	return &ratio{first: first, second: second}
}

type mColor struct {
	color color.Color
}

func (c *mColor) Value() string {
	r, g, b, a := c.color.RGBA()
	return fmt.Sprintf("rgba(%d %d %d %d)", r, g, b, a)
}

func Color(c color.Color) *mColor {
	return &mColor{color: c}
}

type rgba struct {
	r uint32
	g uint32
	b uint32
	a int
}

func (r *rgba) Value() string {
	if r.a >= 0 {
		return fmt.Sprintf("rgba(%d %d %d %d)", r.r, r.g, r.b, r.a)
	}
	return fmt.Sprintf("rgba(%d %d %d)", r.r, r.g, r.b)
}

func Rgb(r, g, b uint32) *rgba {
	return &rgba{r, g, b, -1}
}

func Rgba(r, g, b, a uint32) *rgba {
	return &rgba{r, g, b, int(a)}
}

type hexcolor struct {
	value string
}

func (h *hexcolor) Value() string { return h.value }
func Hex(hex string) *hexcolor    { return &hexcolor{value: hex} }

type hsla struct {
	h uint32
	s uint32
	l uint32
	a int
}

func (h *hsla) Value() string {
	if h.a >= 0 {
		return fmt.Sprintf("hsla(%d %d %d / %d)", h.h, h.s, h.l, h.a)
	}

	return fmt.Sprintf("hsla(%d %d %d)", h.h, h.s, h.l)
}

func Hsl(h, s, l uint32) *hsla     { return &hsla{h, s, l, -1} }
func Hsla(h, s, l, a uint32) *hsla { return &hsla{h, s, l, int(a)} }

type cssColor string

func (c *cssColor) Value() string {
	color := string(*c)
	return color
}

func CSSColor(statement string) *cssColor {
	color := cssColor(statement)
	return &color
}

type withSpacing struct {
	value int
}

func (s *withSpacing) Value() string {
	return fmt.Sprintf("calc(var(--spacing) * %d)", s.value)
}

func NumberWithSpacing(number int) *withSpacing {
	return &withSpacing{value: number}
}
