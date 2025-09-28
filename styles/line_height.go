package styles

import "fmt"

func LeadingNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(lineHeightProp): "1",
		}
	}
}

func Leading(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(lineHeightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func LeadingCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(lineHeightProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func LeadingValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(lineHeightProp): value.String(),
		}
	}
}
