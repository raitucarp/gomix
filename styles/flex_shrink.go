package styles

import "github.com/raitucarp/gomix/value"

func Shrink() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): "1",
		}
	}
}

func ShrinkBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): val.Value(),
		}
	}
}
