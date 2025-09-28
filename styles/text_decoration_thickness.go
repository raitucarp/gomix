package styles

import "fmt"

func Decoration(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorThicknessProp): fmt.Sprintf("%dpx", number),
		}
	}
}

func DecorationFromFont() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorThicknessProp): "from-font",
		}
	}
}

func DecorationAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorThicknessProp): "auto",
		}
	}
}

func DecorationLengthCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorThicknessProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func DecorationValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorThicknessProp): value.String(),
		}
	}
}
