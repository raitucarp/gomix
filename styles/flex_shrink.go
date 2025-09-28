package styles

import (
	"fmt"
	"strconv"
)

func Shrink() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): "1",
		}
	}
}

func ShrinkNumber(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): strconv.Itoa(number),
		}
	}
}

func ShrinkValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): value.String(),
		}
	}
}

func ShrinkCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}
