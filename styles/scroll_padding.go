package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func ScrollP(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingProp): prop,
		}
	}
}

func ScrollPx(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineProp): prop,
		}
	}
}

func ScrollPy(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingBlockProp): prop,
		}
	}
}

func ScrollPs(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineStartProp): prop,
		}
	}
}

func ScrollPe(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingInlineEndProp): prop,
		}
	}
}

func ScrollPt(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingTopProp): prop,
		}
	}
}

func ScrollPr(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingRightProp): prop,
		}
	}
}

func ScrollPb(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingBottomProp): prop,
		}
	}
}

func ScrollPl(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %v)", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(scrollPaddingLeftProp): prop,
		}
	}
}
