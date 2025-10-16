package styles

import (
	"fmt"
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func Order(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(orderProp): strconv.Itoa(number),
		}
	}
}

func NegOrder(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(orderProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func OrderFirst() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(orderProp): "-9999",
		}
	}
}

func OrderLast() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(orderProp): "9999",
		}
	}
}

func OrderBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(orderProp): val.Value(),
		}
	}
}
