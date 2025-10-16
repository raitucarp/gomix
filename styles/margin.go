package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func M(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegM(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): "auto",
		}
	}
}

func MPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): "1px",
		}
	}
}

func NegMPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): "-1px",
		}
	}
}

func MBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginProp): val.Value(),
		}
	}
}

func Mx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMx(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MxAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): "auto",
		}
	}
}

func MxPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): "1px",
		}
	}
}

func NegMxPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): "-1px",
		}
	}
}

func MxBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineProp): val.Value(),
		}
	}
}

func My(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMy(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MyAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): "auto",
		}
	}
}

func MyPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): "1px",
		}
	}
}

func NegMyPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): "-1px",
		}
	}
}

func MyBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBlockProp): val.Value(),
		}
	}
}

func Ms(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMs(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MsAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): "auto",
		}
	}
}

func MsPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): "1px",
		}
	}
}

func NegMsPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): "-1px",
		}
	}
}

func MsBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineStartProp): val.Value(),
		}
	}
}

func Me(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMe(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MeAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): "auto",
		}
	}
}

func MePx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): "1px",
		}
	}
}

func NegMePx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): "-1px",
		}
	}
}

func MeBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginInlineEndProp): val.Value(),
		}
	}
}

func Mt(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMt(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MtAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): "auto",
		}
	}
}

func MtPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): "1px",
		}
	}
}

func NegMtPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): "-1px",
		}
	}
}

func MtBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginTopProp): val.Value(),
		}
	}
}

func Mr(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMr(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MrAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): "auto",
		}
	}
}

func MrPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): "1px",
		}
	}
}

func NegMrPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): "-1px",
		}
	}
}

func MrBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginRightProp): val.Value(),
		}
	}
}

func Mb(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMb(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MbAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): "auto",
		}
	}
}

func MbPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): "1px",
		}
	}
}

func NegMbPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): "-1px",
		}
	}
}

func MbBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginBottomProp): val.Value(),
		}
	}
}

func Ml(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegMl(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func MlAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): "auto",
		}
	}
}

func MlPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): "1px",
		}
	}
}

func NegMlPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): "-1px",
		}
	}
}

func MlBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(marginLeftProp): val.Value(),
		}
	}
}

// TODO: Space-X and Space-Y
