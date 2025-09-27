package styles

import (
	"fmt"
	"strconv"
)

func RowSpan(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): fmt.Sprintf("span %d / span %d", number, number),
		}
	}
}

func RowSpanFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): "1 / -1",
		}
	}
}

func RowSpanCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): fmt.Sprintf("span var(--%s) / span var(--%s)", customProperty, customProperty),
		}
	}
}

func RowSpanValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): fmt.Sprintf("span %v / span %v", value.String(), value.String()),
		}
	}
}

func RowStart(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowStartProp): strconv.Itoa(number),
		}
	}
}

func NegRowStart(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowStartProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowStartAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowStartProp): "auto",
		}
	}
}

func RowStartCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowStartProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func RowStartValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowStartProp): value.String(),
		}
	}
}

func RowEnd(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowEndProp): strconv.Itoa(number),
		}
	}
}

func NegRowEnd(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowEndProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowEndAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowEndProp): "auto",
		}
	}
}

func RowEndCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowEndProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func RowEndValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowEndProp): value.String(),
		}
	}
}

func RowAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): "auto",
		}
	}
}

func Row(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): strconv.Itoa(number),
		}
	}
}

func NegRow(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func RowCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func RowValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridRowProp): value.String(),
		}
	}
}
