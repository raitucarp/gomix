package styles

import "fmt"

func ScrollP(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingProp): prop,
		}
	}
}

func ScrollPx(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineProp): prop,
		}
	}
}

func ScrollPy(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingBlockProp): prop,
		}
	}
}

func ScrollPs(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineStartProp): prop,
		}
	}
}

func ScrollPe(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineEndProp): prop,
		}
	}
}

func ScrollPt(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingTopProp): prop,
		}
	}
}

func ScrollPr(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingRightProp): prop,
		}
	}
}

func ScrollPb(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingBottomProp): prop,
		}
	}
}

func ScrollPl(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingLeftProp): prop,
		}
	}
}
