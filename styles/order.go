package styles

import (
	"fmt"
	"strconv"
)

func Order(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): strconv.Itoa(number),
		}
	}
}

func NegOrder(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func OrderFirst() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): "-9999",
		}
	}
}

func OrderLast() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): "9999",
		}
	}
}

func OrderCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func OrderValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(orderProp): value.String(),
		}
	}
}
