package styles

import "fmt"

func Stroke(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%v", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(positionProp): prop,
		}
	}
}
