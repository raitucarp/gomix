package styles

import (
	"fmt"
)

func Skew(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int:
			propValue = fmt.Sprintf("skewX(%ddeg) skewY(%ddeg)", v, v)
		case float32, float64:
			propValue = fmt.Sprintf("skewX(%.2fdeg) skewY(%.2fdeg)", v, v)
		case customValue:
			propValue = fmt.Sprintf("skewX(%s) skewY(%s)", v.Value(), v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func SkewX(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("skewX(%ddeg)", v)

		case customValue:
			propValue = fmt.Sprintf("skewX(%s)", v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}

func SkewY(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("skewY(%ddeg)", v)

		case customValue:
			propValue = fmt.Sprintf("skewY(%s)", v.Value())
		}

		return &properties{
			string(transformProp): propValue,
		}
	}
}
