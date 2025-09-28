package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func TextInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "inherit",
		}
	}
}

func TextCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "currentColor",
		}
	}
}

func TextTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "transparent",
		}
	}
}

func TextBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func TextWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func TextRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func TextOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func TextAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func TextYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func TextLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func TextGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func TextEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func TextTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func TextCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func TextSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func TextBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func TextIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func TextViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func TextPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func TextFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func TextPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func TextRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func TextSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func TextGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func TextZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func TextNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func TextStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func TextColorCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func TextColorValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): value.String(),
		}
	}
}
