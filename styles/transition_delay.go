package styles

import "fmt"

func Delay(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%dms", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(transitionDelayProp): prop,
		}
	}
}
