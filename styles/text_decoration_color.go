package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func DecorInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): "inherit",
		}
	}
}

func DecorCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): "currentColor",
		}
	}
}

func DecorTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): "transparent",
		}
	}
}

func DecorBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, "black"),
		}
	}
}

func DecorWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, "white"),
		}
	}
}

func DecorRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func DecorOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func DecorAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func DecorYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func DecorLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func DecorGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func DecorEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func DecorTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func DecorCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func DecorSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func DecorBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func DecorIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func DecorViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func DecorPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func DecorFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func DecorPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func DecorRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func DecorSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func DecorGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func DecorZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func DecorNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func DecorStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): s.theme.UseVarKey(themes.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func DecorColorCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func DecorColorValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorColorProp): value.String(),
		}
	}
}
