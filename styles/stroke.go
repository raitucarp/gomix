package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func StrokeInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): "inherit",
		}
	}
}

func StrokeCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): "currentColor",
		}
	}
}

func StrokeTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): "transparent",
		}
	}
}

func StrokeBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func StrokeWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func StrokeRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func StrokeOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func StrokeAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func StrokeYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func StrokeLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func StrokeGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func StrokeEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func StrokeTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func StrokeCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func StrokeSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func StrokeBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func StrokeIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func StrokeViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func StrokePurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func StrokeFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func StrokePink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func StrokeRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func StrokeSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func StrokeGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func StrokeZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func StrokeNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func StrokeStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func StrokeColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(strokeProp): value.Value(),
		}
	}
}
