package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Skew(val any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			propValue = fmt.Sprintf("skewX(%ddeg) skewY(%ddeg)", v, v)
		case float32, float64:
			propValue = fmt.Sprintf("skewX(%#vdeg) skewY(%#vdeg)", v, v)
		case value.Value:
			propValue = fmt.Sprintf("skewX(%s) skewY(%s)", v.Value(), v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func SkewX(val any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("skewX(%ddeg)", v)

		case value.Value:
			propValue = fmt.Sprintf("skewX(%s)", v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func SkewY(val any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("skewY(%ddeg)", v)

		case value.Value:
			propValue = fmt.Sprintf("skewY(%s)", v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}
