package styles

import "github.com/raitucarp/gomix/value"

func Shrink() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexShrinkProp): "1",
		}
	}
}

func ShrinkBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexShrinkProp): val.Value(),
		}
	}
}
