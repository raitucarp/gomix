package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func ScrollM(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginProp): prop,
		}
	}
}

func ScrollMx(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginInlineProp): prop,
		}
	}
}

func ScrollMy(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginBlockProp): prop,
		}
	}
}

func ScrollMs(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginInlineStartProp): prop,
		}
	}
}

func ScrollMe(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginInlineEndProp): prop,
		}
	}
}

func ScrollMt(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginTopProp): prop,
		}
	}
}

func ScrollMr(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginRightProp): prop,
		}
	}
}

func ScrollMb(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginBottomProp): prop,
		}
	}
}

func ScrollMl(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(scrollMarginLeftProp): prop,
		}
	}
}
