package styles

import "fmt"

func ScrollM(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginProp): prop,
		}
	}
}

func ScrollMx(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineProp): prop,
		}
	}
}

func ScrollMy(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginBlockProp): prop,
		}
	}
}

func ScrollMs(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineStartProp): prop,
		}
	}
}

func ScrollMe(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineEndProp): prop,
		}
	}
}

func ScrollMt(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginTopProp): prop,
		}
	}
}

func ScrollMr(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginRightProp): prop,
		}
	}
}

func ScrollMb(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginBottomProp): prop,
		}
	}
}

func ScrollMl(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case customValue:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginLeftProp): prop,
		}
	}
}
