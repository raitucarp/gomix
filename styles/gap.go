package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Gap(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): val.Value(),
		}
	}
}

func GapX(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapXBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): val.Value(),
		}
	}
}

func GapY(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapYBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): val.Value(),
		}
	}
}
