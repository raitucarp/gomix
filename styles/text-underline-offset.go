package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func UnderlineOffset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): fmt.Sprintf("%dpx", number),
		}
	}
}

func NegUnderlineOffset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): fmt.Sprintf("calc(%dpx * -1)", number),
		}
	}
}

func UnderlineOffsetAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): "auto",
		}
	}
}

func UnderlineOffsetBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textUnderlineOffsetProp): val.Value(),
		}
	}
}
