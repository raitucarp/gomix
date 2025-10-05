package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func OutlineInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "inherit",
		}
	}
}

func OutlineCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "currentColor",
		}
	}
}

func OutlineTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "transparent",
		}
	}
}

func OutlineBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func OutlineWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func OutlineRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func OutlineOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func OutlineAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func OutlineYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func OutlineLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func OutlineGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func OutlineEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func OutlineTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func OutlineCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func OutlineSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func OutlineBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func OutlineIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func OutlineViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func OutlinePurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func OutlineFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func OutlinePink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func OutlineRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func OutlineSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func OutlineGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func OutlineZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func OutlineNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func OutlineStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func OutlineColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): value.Value(),
		}
	}
}
