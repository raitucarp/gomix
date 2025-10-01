package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func DropShadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "2xs")),
		}
	}
}

func DropShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "xs")),
		}
	}
}

func DropShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "sm")),
		}
	}
}

func DropShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "md")),
		}
	}
}

func DropShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "lg")),
		}
	}
}

func DropShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "xl")),
		}
	}
}

func DropShadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", s.theme.UseVarKey(themes.Shadow, "2xl")),
		}
	}
}

func DropShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): "drop-shadow(0 0 #0000)",
		}
	}
}

func DropShadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("drop-shadow(%s)", value.Value()),
		}
	}
}

func DropShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): "inherit",
		}
	}
}

func DropShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): "currentColor",
		}
	}
}

func DropShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): "transparent",
		}
	}
}

func DropShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func DropShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func DropShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func DropShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func DropShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func DropShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func DropShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func DropShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func DropShadowEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func DropShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func DropShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func DropShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func DropShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func DropShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func DropShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func DropShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func DropShadowFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func DropShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func DropShadowRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func DropShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func DropShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func DropShadowZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func DropShadowNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func DropShadowStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func DropShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(dropShadowColor): value.Value(),
		}
	}
}
