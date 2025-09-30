package styles

import (
	"fmt"
	"strconv"
)

func FlexNumber(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): strconv.Itoa(number),
		}
	}
}

func FlexFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func FlexAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): "auto",
		}
	}
}

func FlexInitial() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): "0 auto",
		}
	}
}

func FlexNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): "none",
		}
	}
}

func FlexCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func FlexValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexProp): value.String(),
		}
	}
}
