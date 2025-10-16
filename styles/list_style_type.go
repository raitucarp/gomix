package styles

import "github.com/raitucarp/gomix/value"

func ListDisc() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStyleTypeProp): "disc",
		}
	}
}

func ListDecimal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStyleTypeProp): "decimal",
		}
	}
}

func ListNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStyleTypeProp): "none",
		}
	}
}

func List(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStyleTypeProp): val.Value(),
		}
	}
}
