package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Gap(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapProp): val.Value(),
		}
	}
}

func GapX(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapXProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapXBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapXProp): val.Value(),
		}
	}
}

func GapY(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapYProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapYBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gapYProp): val.Value(),
		}
	}
}
