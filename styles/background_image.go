package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Bg(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): value.Value(),
		}
	}
}

func BgNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): "none",
		}
	}
}

func BgLinearToT() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToTr() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToR() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToBr() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom right, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToB() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToBl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to bottom left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToL() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinearToTl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundImageProp): fmt.Sprintf("linear-gradient(to top left, %s)", s.theme.UseVarKey(theme.Custom, "gradient-stops")),
		}
	}
}

func BgLinear(value customValue) ApplyProp {
	return func(s *style) styleProp {
		gradient := ""
		switch v := value.(type) {
		case *angle:
			gradient = fmt.Sprintf("linear-gradient(%s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case *val, *customVariableProp:
			gradient = fmt.Sprintf("linear-gradient(%s, %s))", s.theme.UseVarKey(theme.Custom, "gradient-stops"), value.Value())
		}
		return &properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func BgRadial(value customValue) ApplyProp {
	return func(s *style) styleProp {
		gradient := ""
		switch v := value.(type) {
		case *angle:
			gradient = fmt.Sprintf("radial-gradient(%s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case *val, *customVariableProp:
			gradient = fmt.Sprintf("radial-gradient(%s, %s))", s.theme.UseVarKey(theme.Custom, "gradient-stops"), value.Value())
		}
		return &properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func BgConic(value customValue) ApplyProp {
	return func(s *style) styleProp {
		gradient := ""
		switch v := value.(type) {
		case *angle:
			gradient = fmt.Sprintf("conic-gradient(from %s in oklab, %s)", v.Value(), s.theme.UseVarKey(theme.Custom, "gradient-stops"))
		case *val, *customVariableProp:
			gradient = v.Value()
		}
		return &properties{
			string(backgroundImageProp): gradient,
		}
	}
}

func From(value customValue) ApplyProp {
	propValue := ""
	propName := gradientFrom
	switch v := value.(type) {
	case *mColor, *rgba, *hexcolor, *hsla:
		propValue = v.Value()
	case *percentage:
		propValue = v.Value()
		propName = gradientFromPosition
	case *val, *customVariableProp:
		propValue = v.Value()
	}

	return func(s *style) styleProp {
		return &properties{
			string(propName): propValue,
		}
	}
}

func Via(value customValue) ApplyProp {
	propValue := ""
	propName := gradientVia
	switch v := value.(type) {
	case *mColor, *rgba, *hexcolor, *hsla:
		propValue = v.Value()
	case *percentage:
		propValue = v.Value()
		propName = gradientViaPosition
	case *val, *customVariableProp:
		propValue = v.Value()
	}

	return func(s *style) styleProp {
		return &properties{
			string(propName): propValue,
		}
	}
}

func To(value customValue) ApplyProp {
	propValue := ""
	propName := gradientTo
	switch v := value.(type) {
	case *mColor, *rgba, *hexcolor, *hsla:
		propValue = v.Value()
	case *percentage:
		propValue = v.Value()
		propName = gradientToPosition
	case *val, *customVariableProp:
		propValue = v.Value()
	}

	return func(s *style) styleProp {
		return &properties{
			string(propName): propValue,
		}
	}
}
