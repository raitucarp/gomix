package styles

import "fmt"

func UnderlineOffset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): fmt.Sprintf("%dpx", number),
		}
	}
}

func NegUnderlineOffset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): fmt.Sprintf("calc(%dpx * -1)", number),
		}
	}
}

func UnderlineOffsetAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): "auto",
		}
	}
}

func UnderlineOffsetBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): value.Value(),
		}
	}
}
