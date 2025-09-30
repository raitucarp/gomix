package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func BorderInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "inherit",
		}
	}
}

func BorderCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "currentColor",
		}
	}
}

func BorderTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "transparent",
		}
	}
}

func BorderBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): value.Value(),
		}
	}
}

//--

func BorderXInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "inherit",
		}
	}
}

func BorderXCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "currentColor",
		}
	}
}

func BorderXTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "transparent",
		}
	}
}

func BorderXBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderXWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderXRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderXOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderXAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderXYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderXLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderXGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderXEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderXTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderXCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderXSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderXBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderXIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderXViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderXPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderXFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderXPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderXRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderXSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderXGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderXZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderXNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderXStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderXColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): value.Value(),
		}
	}
}

// --

func BorderYInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "inherit",
		}
	}
}

func BorderYCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "currentColor",
		}
	}
}

func BorderYTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "transparent",
		}
	}
}

func BorderYBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderYWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderYRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderYOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderYAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderYYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderYLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderYGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderYEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderYTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderYCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderYSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderYBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderYIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderYViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderYPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderYFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderYPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderYRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderYSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderYGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderYZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderYNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderYStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderYColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): value.Value(),
		}
	}
}

// --

func BorderStartInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "inherit",
		}
	}
}

func BorderStartCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "currentColor",
		}
	}
}

func BorderStartTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "transparent",
		}
	}
}

func BorderStartBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderStartWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderStartRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderStartOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderStartAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderStartYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderStartLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderStartGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderStartEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderStartTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderStartCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderStartSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderStartBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderStartIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderStartViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderStartPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderStartFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderStartPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderStartRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderStartSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderStartGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderStartZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderStartNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderStartStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderStartColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): value.Value(),
		}
	}
}

// --

func BorderEndInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "inherit",
		}
	}
}

func BorderEndCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "currentColor",
		}
	}
}

func BorderEndTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "transparent",
		}
	}
}

func BorderEndBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderEndWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderEndRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderEndOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderEndAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderEndYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderEndLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderEndGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderEndEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderEndTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderEndCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderEndSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderEndBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderEndIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderEndViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderEndPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderEndFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderEndPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderEndRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderEndSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderEndGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderEndZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderEndNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderEndStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderEndColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): value.Value(),
		}
	}
}

// --

func BorderTopInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "inherit",
		}
	}
}

func BorderTopCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "currentColor",
		}
	}
}

func BorderTopTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "transparent",
		}
	}
}

func BorderTopBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderTopWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderTopRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderTopOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderTopAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderTopYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderTopLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderTopGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderTopEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderTopTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderTopCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderTopSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderTopBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderTopIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderTopViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderTopPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderTopFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderTopPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderTopRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderTopSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderTopGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderTopZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderTopNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderTopStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderTopColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): value.Value(),
		}
	}
}

// --

func BorderRightInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "inherit",
		}
	}
}

func BorderRightCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "currentColor",
		}
	}
}

func BorderRightTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "transparent",
		}
	}
}

func BorderRightBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderRightWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderRightRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderRightOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderRightAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderRightYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderRightLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderRightGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderRightEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderRightTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderRightCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderRightSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderRightBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderRightIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderRightViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderRightPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderRightFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderRightPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderRightRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderRightSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderRightGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderRightZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderRightNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderRightStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderRightColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): value.Value(),
		}
	}
}

// --

func BorderBottomInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "inherit",
		}
	}
}

func BorderBottomCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "currentColor",
		}
	}
}

func BorderBottomTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "transparent",
		}
	}
}

func BorderBottomBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderBottomWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderBottomRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderBottomOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderBottomAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderBottomYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderBottomLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderBottomGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderBottomEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderBottomTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderBottomCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderBottomSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderBottomBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderBottomIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderBottomViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderBottomPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderBottomFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderBottomPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderBottomRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderBottomSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderBottomGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderBottomZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderBottomNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderBottomStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderBottomColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): value.Value(),
		}
	}
}

// --

func BorderLeftInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "inherit",
		}
	}
}

func BorderLeftCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "currentColor",
		}
	}
}

func BorderLeftTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "transparent",
		}
	}
}

func BorderLeftBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BorderLeftWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BorderLeftRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BorderLeftOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BorderLeftAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BorderLeftYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BorderLeftLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BorderLeftGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BorderLeftEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BorderLeftTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BorderLeftCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BorderLeftSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BorderLeftBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BorderLeftIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BorderLeftViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BorderLeftPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BorderLeftFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BorderLeftPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BorderLeftRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BorderLeftSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BorderLeftGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BorderLeftZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BorderLeftNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BorderLeftStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BorderLeftColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): value.Value(),
		}
	}
}

//--

func DivideInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "inherit",
		}
	}
}

func DivideCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "currentColor",
		}
	}
}

func DivideTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "transparent",
		}
	}
}

func DivideBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func DivideWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func DivideRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func DivideOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func DivideAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func DivideYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func DivideLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func DivideGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func DivideEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func DivideTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func DivideCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func DivideSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func DivideBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func DivideIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func DivideViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func DividePurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func DivideFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func DividePink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func DivideRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func DivideSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func DivideGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func DivideZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func DivideNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func DivideStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func DivideColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): value.Value(),
		}
	}
}
