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

func ColSpanBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): fmt.Sprintf("span %v / span %v", value.Value(), value.Value()),
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

func ColStartBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnStartProp): value.Value(),
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

func ColEndBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnEndProp): value.Value(),
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

func ColBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): value.Value(),
		}
	}
}
