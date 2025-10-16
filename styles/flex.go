package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func FlexFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func FlexAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexProp): "auto",
		}
	}
}

func FlexInitial() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexProp): "0 auto",
		}
	}
}

func FlexNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexProp): "none",
		}
	}
}

func FlexBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexProp): val.Value(),
		}
	}
}
