package value

import "fmt"

type Unit[T valConstraint] struct {
	value  T
	suffix string
}

func (u *Unit[T]) Value() string {
	return fmt.Sprintf("%#v%s", u.value, u.suffix)
}

func (u *Unit[T]) LiteralValue() T {
	return u.value
}

func Rem[T valConstraint](value T) *Unit[T] {
	return &Unit[T]{value: value, suffix: "rem"}
}

func Em[T valConstraint](value T) *Unit[T] {
	return &Unit[T]{value: value, suffix: "em"}
}

func Px[T valConstraint](value T) *Unit[T] {
	return &Unit[T]{value: value, suffix: "px"}
}

type PercentageValue struct {
	value float64
}

func (p *PercentageValue) Value() string {
	return fmt.Sprintf("%#v%%", p.value)
}

func Percent(value float64) *PercentageValue {
	return &PercentageValue{value: value}
}

func Pi() *Unit[string] {
	return &Unit[string]{value: "pi", suffix: ""}
}

func E() *Unit[string] { return &Unit[string]{value: "E", suffix: ""} }

func Infinity() *Unit[string]         { return &Unit[string]{value: "infinity", suffix: ""} }
func NegativeInfinity() *Unit[string] { return &Unit[string]{value: "-infinity", suffix: ""} }
func NaN() *Unit[string]              { return &Unit[string]{value: "NaN", suffix: ""} }

type absoluteSize struct{ value string }

func XXSmall() *absoluteSize  { return &absoluteSize{value: "xx-small"} }
func XSmall() *absoluteSize   { return &absoluteSize{value: "x-small"} }
func Small() *absoluteSize    { return &absoluteSize{value: "small"} }
func Medium() *absoluteSize   { return &absoluteSize{value: "medium"} }
func Large() *absoluteSize    { return &absoluteSize{value: "large"} }
func XLarge() *absoluteSize   { return &absoluteSize{value: "x-large"} }
func XXLarge() *absoluteSize  { return &absoluteSize{value: "xx-large"} }
func XXXLarge() *absoluteSize { return &absoluteSize{value: "xxx-large"} }

type Angle struct {
	value float32
	unit  string
}

func (a *Angle) Value() string { return fmt.Sprintf("%.2f%s", a.value, a.unit) }

func Degree(degree float32) *Angle   { return &Angle{value: degree, unit: "deg"} }
func Gradian(gradian float32) *Angle { return &Angle{value: gradian, unit: "grad"} }
func Radian(radian float32) *Angle   { return &Angle{value: radian, unit: "rad"} }
func Turn(turn float32) *Angle       { return &Angle{value: turn, unit: "turn"} }

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

type RgbaColor struct {
	r uint32
	g uint32
	b uint32
	a int
}

func (r *RgbaColor) Value() string {
	if r.a >= 0 {
		return fmt.Sprintf("rgba(%d %d %d %d)", r.r, r.g, r.b, r.a)
	}
	return fmt.Sprintf("rgba(%d %d %d)", r.r, r.g, r.b)
}

func Rgb(r, g, b uint32) *RgbaColor {
	return &RgbaColor{r, g, b, -1}
}

func Rgba(r, g, b, a uint32) *RgbaColor {
	return &RgbaColor{r, g, b, int(a)}
}

type Hexcolor struct {
	value string
}

func (h *Hexcolor) Value() string { return h.value }
func Hex(hex string) *Hexcolor    { return &Hexcolor{value: hex} }

type HslaColor struct {
	h uint32
	s uint32
	l uint32
	a int
}

func (h *HslaColor) Value() string {
	if h.a >= 0 {
		return fmt.Sprintf("hsla(%d %d %d / %d)", h.h, h.s, h.l, h.a)
	}

	return fmt.Sprintf("hsla(%d %d %d)", h.h, h.s, h.l)
}

func Hsl(h, s, l uint32) *HslaColor     { return &HslaColor{h, s, l, -1} }
func Hsla(h, s, l, a uint32) *HslaColor { return &HslaColor{h, s, l, int(a)} }
