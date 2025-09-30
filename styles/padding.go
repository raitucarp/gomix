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

func PBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingProp): value.Value(),
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

func PxBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineProp): value.Value(),
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

func PyBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBlockProp): value.Value(),
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

func PsBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineStartProp): value.Value(),
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

func PeBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingInlineEndProp): value.Value(),
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

func PtBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingTopProp): value.Value(),
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

func PrBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingRightProp): value.Value(),
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

func PbBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingBottomProp): value.Value(),
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

func PlBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(paddingLeftProp): value.Value(),
		}
	}
}
