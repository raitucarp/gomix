package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Decoration(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorThicknessProp): fmt.Sprintf("%dpx", number),
		}
	}
}

func DecorationFromFont() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorThicknessProp): "from-font",
		}
	}
}

func DecorationAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorThicknessProp): "auto",
		}
	}
}

func DecorationBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorThicknessProp): val.Value(),
		}
	}
}
