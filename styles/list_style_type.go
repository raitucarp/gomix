package styles

import "github.com/raitucarp/gomix/value"

func ListDisc() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "disc",
		}
	}
}

func ListDecimal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "decimal",
		}
	}
}

func ListNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "none",
		}
	}
}

func List(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): val.Value(),
		}
	}
}
