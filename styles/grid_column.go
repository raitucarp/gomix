package styles

import (
	"fmt"
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func ColSpan(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): fmt.Sprintf("span %d / span %d", number, number),
		}
	}
}

func ColSpanFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): "1 / -1",
		}
	}
}

func ColSpanBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): fmt.Sprintf("span %v / span %v", val.Value(), val.Value()),
		}
	}
}

func ColStart(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnStartProp): strconv.Itoa(number),
		}
	}
}

func NegColStart(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnStartProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColStartAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnStartProp): "auto",
		}
	}
}

func ColStartBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnStartProp): val.Value(),
		}
	}
}

func ColEnd(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnEndProp): strconv.Itoa(number),
		}
	}
}

func NegColEnd(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnEndProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColEndAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnEndProp): "auto",
		}
	}
}

func ColEndBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnEndProp): val.Value(),
		}
	}
}

func ColAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): "auto",
		}
	}
}

func Col(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): strconv.Itoa(number),
		}
	}
}

func NegCol(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridColumnProp): val.Value(),
		}
	}
}
