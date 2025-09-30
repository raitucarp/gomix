package styles

import "fmt"

func outlineOffsetValue(value ...customValue) string {
	propValue := ""
	if len(value) == 0 {
		propValue = "1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				if n <= 0 {

					propValue = fmt.Sprintf("%dpx", n)
				} else {
					propValue = fmt.Sprintf("calc(%dpx * -1)", n)
				}
			default:
				propValue = v.Value()
			}
		case *customVariableProp:
			propValue = v.Value()
		}
	}

	return propValue
}

func OutlineOffset(value ...customValue) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(outlineOffsetProp): outlineOffsetValue(value...),
		}
	}
}
