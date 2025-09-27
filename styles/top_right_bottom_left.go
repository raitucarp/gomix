package styles

import (
	"fmt"
	"strconv"
)

func Inset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInset(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegInsetFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func InsetPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): "1px",
		}
	}
}

func NegInsetPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): "-1px",
		}
	}
}

func InsetFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): "100%%",
		}
	}
}

func NegInsetFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): "-100%%",
		}
	}
}

func InsetAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): "auto",
		}
	}
}

func InsetCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func InsetValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetProp): value.String(),
		}
	}
}

//--

func InsetX(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInsetX(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetXFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegInsetXFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func InsetXPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): "1px",
		}
	}
}

func NegInsetXPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): "-1px",
		}
	}
}

func InsetXFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): "100%%",
		}
	}
}

func NegInsetXFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): "-100%%",
		}
	}
}

func InsetXAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): "auto",
		}
	}
}

func InsetXCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func InsetXValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetInlineProp): value.String(),
		}
	}
}

//--

func InsetY(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegInsetY(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func InsetYFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegInsetYFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func InsetYPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): "1px",
		}
	}
}

func NegInsetYPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): "-1px",
		}
	}
}

func InsetYFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): "100%%",
		}
	}
}

func NegInsetYFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): "-100%%",
		}
	}
}

func InsetYAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): "auto",
		}
	}
}

func InsetYCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func InsetYValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetBlockProp): value.String(),
		}
	}
}

//--

func Start(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegStart(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func StartFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegStartFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func StartPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): "1px",
		}
	}
}

func NegStartPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): "-1px",
		}
	}
}

func StartFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): "100%%",
		}
	}
}

func NegStartFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): "-100%%",
		}
	}
}

func StartAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): "auto",
		}
	}
}

func StartCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func StartValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetStartProp): value.String(),
		}
	}
}

//--

func End(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegEnd(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func EndFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegEndFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func EndPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): "1px",
		}
	}
}

func NegEndPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): "-1px",
		}
	}
}

func EndFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): "100%%",
		}
	}
}

func NegEndFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): "-100%%",
		}
	}
}

func EndAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): "auto",
		}
	}
}

func EndCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func EndValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetEndProp): value.String(),
		}
	}
}

//--

func Top(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegTop(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func TopFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegTopFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func TopPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): "1px",
		}
	}
}

func NegTopPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): "-1px",
		}
	}
}

func TopFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): "100%%",
		}
	}
}

func NegTopFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): "-100%%",
		}
	}
}

func TopAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): "auto",
		}
	}
}

func TopCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func TopValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(topProp): value.String(),
		}
	}
}

//--

func Right(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegRight(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func RightFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegRightFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func RightPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): "1px",
		}
	}
}

func NegRightPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): "-1px",
		}
	}
}

func RightFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): "100%%",
		}
	}
}

func NegRightFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): "-100%%",
		}
	}
}

func RightAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): "auto",
		}
	}
}

func RightCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func RightValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(rightProp): value.String(),
		}
	}
}

// --

func Bottom(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegBottom(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func BottomFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegBottomFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func BottomPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): "1px",
		}
	}
}

func NegBottomPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): "-1px",
		}
	}
}

func BottomFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): "100%%",
		}
	}
}

func NegBottomFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): "-100%%",
		}
	}
}

func BottomAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): "auto",
		}
	}
}

func BottomCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func BottomValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(bottomProp): value.String(),
		}
	}
}

// --

func Left(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegLeft(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func LeftFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): fmt.Sprintf("calc(%s * 100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func NegLeftFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): fmt.Sprintf("calc(%s * -100%%)", strconv.FormatFloat(fraction, 'f', -1, 32)),
		}
	}
}

func LeftPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): "1px",
		}
	}
}

func NegLeftPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): "-1px",
		}
	}
}

func LeftFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): "100%%",
		}
	}
}

func NegLeftFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): "-100%%",
		}
	}
}

func LeftAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): "auto",
		}
	}
}

func LeftCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func LeftValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(leftProp): value.String(),
		}
	}
}
