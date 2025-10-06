package styles

import "github.com/raitucarp/gomix/value"

func Content(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): val.Value(),
		}
	}
}

func ContentNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): "none",
		}
	}
}
