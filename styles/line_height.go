package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

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

func LeadingBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(lineHeightProp): val.Value(),
		}
	}
}
