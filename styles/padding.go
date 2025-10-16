package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func P(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingProp): "1px",
		}
	}
}

func PBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingProp): val.Value(),
		}
	}
}

func Px(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PxPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineProp): "1px",
		}
	}
}

func PxBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineProp): val.Value(),
		}
	}
}

func Py(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBlockProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PyPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBlockProp): "1px",
		}
	}
}

func PyBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBlockProp): val.Value(),
		}
	}
}

func Ps(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineStartProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PsPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineStartProp): "1px",
		}
	}
}

func PsBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineStartProp): val.Value(),
		}
	}
}

func Pe(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineEndProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PePx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineEndProp): "1px",
		}
	}
}

func PeBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingInlineEndProp): val.Value(),
		}
	}
}

func Pt(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingTopProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PtPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingTopProp): "1px",
		}
	}
}

func PtBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingTopProp): val.Value(),
		}
	}
}

func Pr(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingRightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PrPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingRightProp): "1px",
		}
	}
}

func PrBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingRightProp): val.Value(),
		}
	}
}

func Pb(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBottomProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PbPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBottomProp): "1px",
		}
	}
}

func PbBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingBottomProp): val.Value(),
		}
	}
}

func Pl(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingLeftProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func PlPx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingLeftProp): "1px",
		}
	}
}

func PlBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(paddingLeftProp): val.Value(),
		}
	}
}
