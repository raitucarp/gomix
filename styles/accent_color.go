package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func AccentInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): "inherit",
		}
	}
}

func AccentCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): "currentColor",
		}
	}
}

func AccentTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): "transparent",
		}
	}
}

func AccentBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func AccentWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func AccentRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func AccentOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func AccentAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func AccentYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func AccentLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func AccentGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func AccentEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func AccentTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func AccentCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func AccentSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func AccentBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func AccentIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func AccentViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func AccentPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func AccentFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func AccentPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func AccentRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func AccentSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func AccentGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func AccentZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func AccentNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func AccentStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func Accent(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(accentColorProp): value.Value(),
		}
	}
}
