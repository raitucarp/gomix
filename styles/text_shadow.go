package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func TextShadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "2xs"),
		}
	}
}

func TextShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "xs"),
		}
	}
}

func TextShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "sm"),
		}
	}
}

func TextShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "md"),
		}
	}
}

func TextShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "lg"),
		}
	}
}

func TextShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "xl"),
		}
	}
}

func TextShadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): s.theme.UseVarKey(themes.Shadow, "2xl"),
		}
	}
}

func TextShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): "0 0 #0000",
		}
	}
}

func TextShadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowProp): value.Value(),
		}
	}
}

func TextShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): "inherit",
		}
	}
}

func TextShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): "currentColor",
		}
	}
}

func TextShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): "transparent",
		}
	}
}

func TextShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func TextShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func TextShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func TextShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func TextShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func TextShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func TextShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func TextShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func TextShadowEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func TextShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func TextShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func TextShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func TextShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func TextShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func TextShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func TextShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func TextShadowFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func TextShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func TextShadowRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func TextShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func TextShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func TextShadowZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func TextShadowNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func TextShadowStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func TextShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textShadowColorProp): value.Value(),
		}
	}
}
