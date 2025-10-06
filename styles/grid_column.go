package styles

import (
	"fmt"
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func ColSpan(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("span %d / span %d", number, number),
		}
	}
}

func ColSpanFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): "1 / -1",
		}
	}
}

func ColSpanBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("span %v / span %v", val.Value(), val.Value()),
		}
	}
}

func ColStart(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): strconv.Itoa(number),
		}
	}
}

func NegColStart(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColStartAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): "auto",
		}
	}
}

func ColStartBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): val.Value(),
		}
	}
}

func ColEnd(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): strconv.Itoa(number),
		}
	}
}

func NegColEnd(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColEndAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): "auto",
		}
	}
}

func ColEndBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): val.Value(),
		}
	}
}

func ColAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): "auto",
		}
	}
}

func Col(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): strconv.Itoa(number),
		}
	}
}

func NegCol(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("calc(%d * -1)", number),
		}
	}
}

func ColBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): val.Value(),
		}
	}
}
