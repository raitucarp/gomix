package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func Bg(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): val.Value(),
		}
	}
}

func BgNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): "none",
		}
	}
}

func BgLinearToT() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToTr() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToR() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToBr() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToB() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToBl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToL() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToTl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinear(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		gradient := ""
		switch v := val.(type) {
		case *value.Angle:
			gradient = fmt.Sprintf("linear-gradient(%s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case value.Value, *value.CustomProperty:
			gradient = fmt.Sprintf("linear-gradient(%s, %s))", s.theme.UseVarKey(theme.Custom, "gradient-stops"), val.Value())
		}
		return &Properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func BgRadial(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		gradient := ""
		switch v := val.(type) {
		case *value.Angle:
			gradient = fmt.Sprintf("radial-gradient(%s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case value.Value, *value.CustomProperty:
			val.Value()
			gradient = fmt.Sprintf("radial-gradient(%s, %s))", s.theme.UseVarKey(theme.Custom, "gradient-stops"), val.Value())
		}
		return &Properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func BgConic(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		gradient := ""
		switch v := val.(type) {
		case *value.Angle:
			gradient = fmt.Sprintf("conic-gradient(from %s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case value.Value, *value.CustomProperty:
			gradient = v.Value()
		}
		return &Properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func From(val value.Value) ApplyProp {
	propValue := ""
	propName := gradientFrom
	switch v := val.(type) {
	case *value.RgbaColor, *value.Hexcolor, *value.HslaColor:
		propValue = v.Value()
	case *value.PercentageValue:
		propValue = v.Value()
		propName = gradientFromPosition
	case value.Value, *value.CustomProperty:
		propValue = v.Value()
	}

	return func(s *Style) StyleProp {
		return &Properties{
			string(propName): propValue,
		}
	}
}

func Via(val value.Value) ApplyProp {
	propValue := ""
	propName := gradientVia
	switch v := val.(type) {
	case *value.RgbaColor, *value.Hexcolor, *value.HslaColor:
		propValue = v.Value()
	case *value.PercentageValue:
		propValue = v.Value()
		propName = gradientViaPosition
	case value.Value, *value.CustomProperty:
		propValue = v.Value()
	}

	return func(s *Style) StyleProp {
		return &Properties{
			string(propName): propValue,
		}
	}
}

func To(val value.Value) ApplyProp {
	propValue := ""
	propName := gradientTo
	switch v := val.(type) {
	case *value.RgbaColor, *value.Hexcolor, *value.HslaColor:
		propValue = v.Value()
	case *value.PercentageValue:
		propValue = v.Value()
		propName = gradientToPosition
	case value.Value, *value.CustomProperty:
		propValue = v.Value()
	}

	return func(s *Style) StyleProp {
		return &Properties{
			string(propName): propValue,
		}
	}
}
