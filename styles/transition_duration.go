package styles

import "fmt"

func Duration(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%dms", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(transitionDurationProp): prop,
		}
	}
}

func DurationInitial() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionDurationProp): "initial",
		}
	}
}
