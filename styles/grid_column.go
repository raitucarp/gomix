package styles

import (
	"fmt"
	"strconv"
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

func ColSpanCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("span var(--%s) / span var(--%s)", customProperty, customProperty),
		}
	}
}

func ColSpanValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("span %v / span %v", value.String(), value.String()),
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

func ColStartCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ColStartValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): value.String(),
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

func ColEndCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ColEndValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): value.String(),
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

func ColCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ColValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): value.String(),
		}
	}
}
