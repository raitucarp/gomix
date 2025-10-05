package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func FillInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): "inherit",
		}
	}
}

func FillCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): "currentColor",
		}
	}
}

func FillTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): "transparent",
		}
	}
}

func FillBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func FillWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func FillRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}

func FillOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}

func FillAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}

func FillYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}

func FillLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}

func FillGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}

func FillEmerald(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("emerald-%d", scale)),
		}
	}
}

func FillTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}

func FillCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}

func FillSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}

func FillBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}

func FillIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}

func FillViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}

func FillPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}

func FillFuchsia(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("fuchsia-%d", scale)),
		}
	}
}

func FillPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}

func FillRose(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("rose-%d", scale)),
		}
	}
}

func FillSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}

func FillGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}

func FillZinc(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("zinc-%d", scale)),
		}
	}
}

func FillNeutral(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("neutral-%d", scale)),
		}
	}
}

func FillStone(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("stone-%d", scale)),
		}
	}
}

func Fill(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fillProp): value.Value(),
		}
	}
}
