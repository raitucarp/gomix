package styles

import "fmt"

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

func ListCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ListValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): value.String(),
		}
	}
}
