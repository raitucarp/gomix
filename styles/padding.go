package styles

import "fmt"

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

func PCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): value.String(),
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

func PxCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PxValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): value.String(),
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

func PyCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PyValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): value.String(),
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

func PsCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PsValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): value.String(),
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

func PeCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PeValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): value.String(),
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

func PtCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PtValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): value.String(),
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

func PrCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PrValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): value.String(),
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

func PbCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PbValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): value.String(),
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

func PlCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func PlValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): value.String(),
		}
	}
}
