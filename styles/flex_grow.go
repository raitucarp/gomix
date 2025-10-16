package styles

import (
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func Grow() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexGrowProp): "1",
		}
	}
}

func GrowNumber(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexGrowProp): strconv.Itoa(number),
		}
	}
}

func GrowValue(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexGrowProp): val.Value(),
		}
	}
}
