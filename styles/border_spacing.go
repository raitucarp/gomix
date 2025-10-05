package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func BorderSpacing(spacing any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %.2f)", v)
		case customValue:
			prop = v.Value()
		}

		return &properties{
			string(borderSpacing): prop,
		}
	}
}

func BorderSpacingX(spacing any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %.2f) %s", v, s.theme.UseVarKey(theme.Custom, "border-spacing-y"))
		case customValue:
			prop = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(theme.Custom, "border-spacing-y"))
		}

		return &properties{
			string(borderSpacing): prop,
		}
	}
}

func BorderSpacingY(spacing any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %.2f) %s", v, s.theme.UseVarKey(theme.Custom, "border-spacing-x"))
		case customValue:
			prop = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(theme.Custom, "border-spacing-x"))
		}

		return &properties{
			string(borderSpacing): prop,
		}
	}
}
