package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Inset(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInset(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegInsetFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func InsetPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): "1px",
		}
	}
}

func NegInsetPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): "-1px",
		}
	}
}

func InsetFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): "100%%",
		}
	}
}

func NegInsetFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): "-100%%",
		}
	}
}

func InsetAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): "auto",
		}
	}
}

func InsetBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetProp): val.Value(),
		}
	}
}

//--

func InsetX(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInsetX(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetXFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegInsetXFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func InsetXPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): "1px",
		}
	}
}

func NegInsetXPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): "-1px",
		}
	}
}

func InsetXFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): "100%%",
		}
	}
}

func NegInsetXFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): "-100%%",
		}
	}
}

func InsetXAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): "auto",
		}
	}
}

func InsetXBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetInlineProp): val.Value(),
		}
	}
}

//--

func InsetY(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInsetY(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetYFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegInsetYFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func InsetYPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): "1px",
		}
	}
}

func NegInsetYPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): "-1px",
		}
	}
}

func InsetYFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): "100%%",
		}
	}
}

func NegInsetYFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): "-100%%",
		}
	}
}

func InsetYAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): "auto",
		}
	}
}

func InsetYBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetBlockProp): val.Value(),
		}
	}
}

//--

func Start(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegStart(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func StartFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegStartFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func StartPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): "1px",
		}
	}
}

func NegStartPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): "-1px",
		}
	}
}

func StartFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): "100%%",
		}
	}
}

func NegStartFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): "-100%%",
		}
	}
}

func StartAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): "auto",
		}
	}
}

func StartBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetStartProp): val.Value(),
		}
	}
}

//--

func End(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegEnd(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func EndFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegEndFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func EndPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): "1px",
		}
	}
}

func NegEndPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): "-1px",
		}
	}
}

func EndFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): "100%%",
		}
	}
}

func NegEndFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): "-100%%",
		}
	}
}

func EndAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): "auto",
		}
	}
}

func EndBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(insetEndProp): val.Value(),
		}
	}
}

//--

func Top(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegTop(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func TopFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegTopFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func TopPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): "1px",
		}
	}
}

func NegTopPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): "-1px",
		}
	}
}

func TopFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): "100%%",
		}
	}
}

func NegTopFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): "-100%%",
		}
	}
}

func TopAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): "auto",
		}
	}
}

func TopBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(topProp): val.Value(),
		}
	}
}

//--

func Right(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegRight(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func RightFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegRightFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func RightPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): "1px",
		}
	}
}

func NegRightPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): "-1px",
		}
	}
}

func RightFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): "100%%",
		}
	}
}

func NegRightFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): "-100%%",
		}
	}
}

func RightAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): "auto",
		}
	}
}

func RightBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(rightProp): val.Value(),
		}
	}
}

// --

func Bottom(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegBottom(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func BottomFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegBottomFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func BottomPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): "1px",
		}
	}
}

func NegBottomPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): "-1px",
		}
	}
}

func BottomFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): "100%%",
		}
	}
}

func NegBottomFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): "-100%%",
		}
	}
}

func BottomAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): "auto",
		}
	}
}

func BottomBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(bottomProp): val.Value(),
		}
	}
}

// --

func Left(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegLeft(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func LeftFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func NegLeftFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): fmt.Sprintf("calc(%#v * -100%%)", fraction),
		}
	}
}

func LeftPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): "1px",
		}
	}
}

func NegLeftPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): "-1px",
		}
	}
}

func LeftFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): "100%%",
		}
	}
}

func NegLeftFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): "-100%%",
		}
	}
}

func LeftAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): "auto",
		}
	}
}

func LeftBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(leftProp): val.Value(),
		}
	}
}
