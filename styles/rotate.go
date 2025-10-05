package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func RotateNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rotateProp): "none",
		}
	}
}

func Rotate(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%ddeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%d", v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2fdeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%.2f", v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2fdeg * -1)", v)
			} else {
				propValue = fmt.Sprintf("%.2f", v)
			}
		case customValue:
			propValue = v.Value()
		}

		return &properties{
			string(rotateProp): propValue,
		}
	}
}

func RotateX(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("rotateX(%vdeg) %s", v, s.theme.UseVarKey(theme.Custom, "rotate-y"))

		case customValue:
			propValue = fmt.Sprintf("rotateX(%s) %s", v.Value(), s.theme.UseVarKey(theme.Custom, "rotate-y"))
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func RotateY(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s rotateY(%ddeg)", s.theme.UseVarKey(theme.Custom, "rotate-x"), v)

		case customValue:
			propValue = fmt.Sprintf("%s rotateY(%s)", s.theme.UseVarKey(theme.Custom, "rotate-x"), v.Value())

		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func RotateZ(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s rotateZ(%ddeg)",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				v)
		case customValue:
			propValue = fmt.Sprintf("%s %s rotateZ(%s)",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				v.Value())

		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}
