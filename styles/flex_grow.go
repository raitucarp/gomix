package styles

import (
	"fmt"
	"strconv"
)

func Grow() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexGrowProp): "1",
		}
	}
}

func GrowNumber(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexGrowProp): strconv.Itoa(number),
		}
	}
}

func GrowValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexGrowProp): value.String(),
		}
	}
}

func GrowCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexGrowProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}
