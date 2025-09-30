package styles

import "fmt"

func opacityValue(value customValue) string {
	propValue := ""
	switch v := value.(type) {
	case *val:
		switch n := v.value.(type) {
		case int:
			propValue = fmt.Sprintf("%d%%", n)
		default:
			propValue = v.Value()
		}
	case *customVariableProp:
		propValue = v.Value()
	}

	return propValue
}

func Opacity(value customValue) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(opacityProp): opacityValue(value),
		}
	}
}
