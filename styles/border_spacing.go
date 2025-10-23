package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func BorderSpacing(spacing any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %#v)", v)
		case value.Value:
			prop = v.Value()
		}

		return &Properties{
			string(borderSpacing): prop,
		}
	}
}

func BorderSpacingX(spacing any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %#v) %s", v, s.theme.UseVarKey(themeCustom, "border-spacing-y"))
		case value.Value:
			prop = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(themeCustom, "border-spacing-y"))
		}

		return &Properties{
			string(borderSpacing): prop,
		}
	}
}

func BorderSpacingY(spacing any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := spacing.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("calc(var(--spacing) * %#v) %s", v, s.theme.UseVarKey(themeCustom, "border-spacing-x"))
		case value.Value:
			prop = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(themeCustom, "border-spacing-x"))
		}

		return &Properties{
			string(borderSpacing): prop,
		}
	}
}
