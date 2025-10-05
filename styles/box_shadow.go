package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Shadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xs"),
		}
	}
}

func ShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xs"),
		}
	}
}

func ShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "sm"),
		}
	}
}

func ShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "md"),
		}
	}
}

func ShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "lg"),
		}
	}
}

func ShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xl"),
		}
	}
}

func Shadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xl"),
		}
	}
}

func ShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): "0 0 #0000",
		}
	}
}

func Shadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): value.Value(),
		}
	}
}

func ShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "inherit",
		}
	}
}

func ShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "currentColor",
		}
	}
}

func ShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "transparent",
		}
	}
}

func ShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func ShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func ShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func ShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func ShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func ShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func ShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func ShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func ShadowEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func ShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func ShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func ShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func ShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func ShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func ShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func ShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func ShadowFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func ShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func ShadowRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func ShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func ShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func ShadowZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func ShadowNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func ShadowStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func ShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): value.Value(),
		}
	}
}

//--

func InsetShadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xs"),
		}
	}
}

func InsetShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xs"),
		}
	}
}

func InsetShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "sm"),
		}
	}
}

func InsetShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "md"),
		}
	}
}

func InsetShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "lg"),
		}
	}
}

func InsetShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xl"),
		}
	}
}

func InsetShadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xl"),
		}
	}
}

func InsetShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): "inset 0 0 #0000",
		}
	}
}

func InsetShadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): value.Value(),
		}
	}
}

func InsetShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "inherit",
		}
	}
}

func InsetShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "currentColor",
		}
	}
}

func InsetShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "transparent",
		}
	}
}

func InsetShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func InsetShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func InsetShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func InsetShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func InsetShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func InsetShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func InsetShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func InsetShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func InsetShadowEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func InsetShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func InsetShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func InsetShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func InsetShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func InsetShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func InsetShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func InsetShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func InsetShadowFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func InsetShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func InsetShadowRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func InsetShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func InsetShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func InsetShadowZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func InsetShadowNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func InsetShadowStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func InsetShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): value.Value(),
		}
	}
}

// --

func Ring(value ...customValue) ApplyProp {
	propValue := ""
	if len(value) == 0 {
		propValue = "0 0 0 1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				propValue = fmt.Sprintf("0 0 0 %dpx", n)
			default:
				propValue = fmt.Sprintf("0 0 0 %s", v.Value())
			}
		case *customVariableProp:
			propValue = fmt.Sprintf("0 0 0 %s", v.Value())
		}
	}

	return func(s *style) styleProp {
		return &properties{
			string(ringShadowProp): propValue,
		}
	}
}

func RingInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "inherit",
		}
	}
}

func RingCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "currentColor",
		}
	}
}

func RingTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "transparent",
		}
	}
}

func RingBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func RingWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func RingRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func RingOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func RingAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func RingYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func RingLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func RingGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func RingEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func RingTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func RingCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func RingSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func RingBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func RingIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func RingViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func RingPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func RingFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func RingPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func RingRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func RingSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func RingGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func RingZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func RingNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func RingStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

//--

func InsetRing(value ...customValue) ApplyProp {
	propValue := ""
	if len(value) == 0 {
		propValue = "0 0 0 1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				propValue = fmt.Sprintf("0 0 0 %dpx", n)
			default:
				propValue = fmt.Sprintf("0 0 0 %s", v.Value())
			}
		case *customVariableProp:
			propValue = fmt.Sprintf("0 0 0 %s", v.Value())
		}
	}

	return func(s *style) styleProp {
		return &properties{
			string(insetRingShadowProp): propValue,
		}
	}
}

func InsetRingInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "inherit",
		}
	}
}

func InsetRingCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "currentColor",
		}
	}
}

func InsetRingTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "transparent",
		}
	}
}

func InsetRingBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func InsetRingWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func InsetRingRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func InsetRingOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func InsetRingAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func InsetRingYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func InsetRingLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func InsetRingGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func InsetRingEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func InsetRingTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func InsetRingCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func InsetRingSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func InsetRingBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func InsetRingIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func InsetRingViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func InsetRingPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func InsetRingFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func InsetRingPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func InsetRingRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func InsetRingSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func InsetRingGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func InsetRingZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func InsetRingNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func InsetRingStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}
