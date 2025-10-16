package styles

import (
	"fmt"
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func RowSpan(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): fmt.Sprintf("span %d / span %d", number, number),
		}
	}
}

func RowSpanFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): "1 / -1",
		}
	}
}

func RowSpanBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): fmt.Sprintf("span %v / span %v", val.Value(), val.Value()),
		}
	}
}

func RowStart(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowStartProp): strconv.Itoa(number),
		}
	}
}

func NegRowStart(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowStartProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowStartAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowStartProp): "auto",
		}
	}
}

func RowStartBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowStartProp): val.Value(),
		}
	}
}

func RowEnd(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowEndProp): strconv.Itoa(number),
		}
	}
}

func NegRowEnd(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowEndProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowEndAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowEndProp): "auto",
		}
	}
}

func RowEndBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowEndProp): val.Value(),
		}
	}
}

func RowAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): "auto",
		}
	}
}

func Row(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): strconv.Itoa(number),
		}
	}
}

func NegRow(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridRowProp): val.Value(),
		}
	}
}
