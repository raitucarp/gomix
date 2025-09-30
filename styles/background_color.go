package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func BgInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): "inherit",
		}
	}
}

func BgCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): "currentColor",
		}
	}
}

func BgTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): "transparent",
		}
	}
}

func BgBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func BgWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func BgRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func BgOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func BgAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func BgYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func BgLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func BgGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func BgEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func BgTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func BgCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func BgSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func BgBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func BgIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func BgViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func BgPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func BgFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func BgPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func BgRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func BgSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func BgGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func BgZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func BgNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func BgStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func BgColor(color customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundColorProp): color.Value(),
		}
	}
}
