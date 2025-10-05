package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func CaretInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): "inherit",
		}
	}
}

func CaretCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): "currentColor",
		}
	}
}

func CaretTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): "transparent",
		}
	}
}

func CaretBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func CaretWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func CaretRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func CaretOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func CaretAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func CaretYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func CaretLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func CaretGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func CaretEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func CaretTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func CaretCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func CaretSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func CaretBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func CaretIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func CaretViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func CaretPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func CaretFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func CaretPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func CaretRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func CaretSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func CaretGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func CaretZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func CaretNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func CaretStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func Caret(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(caretColorProp): value.Value(),
		}
	}
}
