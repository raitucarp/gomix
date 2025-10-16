package styles

import "github.com/raitucarp/gomix/value"

func Content(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(contentProp): val.Value(),
		}
	}
}

func ContentNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(contentProp): "none",
		}
	}
}
