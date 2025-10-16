package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func RotateNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rotateProp): "none",
		}
	}
}

func Rotate(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%ddeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%d", v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#vdeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%#v", v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#vdeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%#v", v)
			}
		case value.Value:
			propValue = v.Value()
		}

		return &Properties{
			string(rotateProp): propValue,
		}
	}
}

func RotateX(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("rotateX(%vdeg) %s", v, s.theme.UseVarKey(theme.Custom, "rotate-y"))

		case value.Value:
			propValue = fmt.Sprintf("rotateX(%s) %s", v.Value(), s.theme.UseVarKey(theme.Custom, "rotate-y"))
		}

		return &Properties{
			string(transformProp): propValue,
		}
	}
}

func RotateY(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s rotateY(%ddeg)", s.theme.UseVarKey(theme.Custom, "rotate-x"), v)

		case value.Value:
			propValue = fmt.Sprintf("%s rotateY(%s)", s.theme.UseVarKey(theme.Custom, "rotate-x"), v.Value())

		}

		return &Properties{
			string(transformProp): propValue,
		}
	}
}

func RotateZ(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s rotateZ(%ddeg)",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				v)
		case value.Value:
			propValue = fmt.Sprintf("%s %s rotateZ(%s)",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				v.Value())

		}

		return &Properties{
			string(transformProp): propValue,
		}
	}
}
