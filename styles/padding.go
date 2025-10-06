package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func P(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): "1px",
		}
	}
}

func PBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): val.Value(),
		}
	}
}

func Px(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PxPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): "1px",
		}
	}
}

func PxBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): val.Value(),
		}
	}
}

func Py(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PyPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): "1px",
		}
	}
}

func PyBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): val.Value(),
		}
	}
}

func Ps(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PsPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): "1px",
		}
	}
}

func PsBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): val.Value(),
		}
	}
}

func Pe(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PePx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): "1px",
		}
	}
}

func PeBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): val.Value(),
		}
	}
}

func Pt(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PtPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): "1px",
		}
	}
}

func PtBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): val.Value(),
		}
	}
}

func Pr(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PrPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): "1px",
		}
	}
}

func PrBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): val.Value(),
		}
	}
}

func Pb(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PbPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): "1px",
		}
	}
}

func PbBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): val.Value(),
		}
	}
}

func Pl(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PlPx(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): "1px",
		}
	}
}

func PlBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): val.Value(),
		}
	}
}
