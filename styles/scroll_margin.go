package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func ScrollM(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginProp): prop,
		}
	}
}

func ScrollMx(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineProp): prop,
		}
	}
}

func ScrollMy(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginBlockProp): prop,
		}
	}
}

func ScrollMs(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineStartProp): prop,
		}
	}
}

func ScrollMe(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginInlineEndProp): prop,
		}
	}
}

func ScrollMt(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginTopProp): prop,
		}
	}
}

func ScrollMr(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginRightProp): prop,
		}
	}
}

func ScrollMb(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginBottomProp): prop,
		}
	}
}

func ScrollMl(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollMarginLeftProp): prop,
		}
	}
}
