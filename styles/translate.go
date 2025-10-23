package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func TranslateNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): "none",
		}
	}
}

func Translate(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * %d) calc(var(--spacing) * %d)", v, v)
		case value.RatioValue:
			propValue = fmt.Sprintf("calc(%s * 100%%) calc(%s * 100%%)", v.Value(), v.Value())
		case value.Value:
			propValue = fmt.Sprintf("%s %s", v, v)
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslate(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * -%d) calc(var(--spacing) * -%d)", v, v)
		case value.RatioValue:
			propValue = fmt.Sprintf("calc(%s * -100%%) calc(%s * -100%%)", v.Value(), v.Value())
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): `100% 100%`,
		}
	}
}

func NegTranslateFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): `-100% -100%`,
		}
	}
}

func TranslatePx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): `1px 1px`,
		}
	}
}

func NegTranslatePx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): `-1px -1px`,
		}
	}
}

func TranslateX(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * %d) %s",
				v,
				s.theme.UseVarKey(themeCustom, "translate-y"),
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("calc(%s * 100%%) %s", v.Value(), s.theme.UseVarKey(themeCustom, "translate-y"))
		case value.Value:
			propValue = fmt.Sprintf("%s %s", v, s.theme.UseVarKey(themeCustom, "translate-y"))
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateX(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * -%d) %s",
				v,
				s.theme.UseVarKey(themeCustom, "translate-y"),
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("calc(%s * -100%%) %s", v.Value(), s.theme.UseVarKey(themeCustom, "translate-y"))
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateXFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("100%% %s",
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}

func NegTranslateXFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("-100%% %s",
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}

func TranslateXPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("1px %s",
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}

func NegTranslateXPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("-1px %s",
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}

func TranslateY(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s calc(var(--spacing) * %d)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				v,
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("%s calc(%s * 100%%)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				v.Value(),
			)
		case value.Value:
			propValue = fmt.Sprintf("%s %s",
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v,
			)
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateY(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s calc(var(--spacing) * -%d)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				v,
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("%s calc(%s * -100%%)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				v.Value(),
			)
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateYFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s 100%%",
				s.theme.UseVarKey(themeCustom, "translate-x"),
			),
		}
	}
}

func NegTranslateYFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s -100%%",
				s.theme.UseVarKey(themeCustom, "translate-x"),
			),
		}
	}
}

func TranslateYPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s 1px",
				s.theme.UseVarKey(themeCustom, "translate-x"),
			),
		}
	}
}

func NegTranslateYPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s -1px",
				s.theme.UseVarKey(themeCustom, "translate-x"),
			),
		}
	}
}

func TranslateZ(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s calc(var(--spacing) * %d)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v,
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("%s %s calc(%s * 100%%)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v.Value(),
			)
		case value.Value:
			propValue = fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v,
			)
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateZ(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s calc(var(--spacing) * -%d)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v,
			)
		case value.RatioValue:
			propValue = fmt.Sprintf("%s %s calc(%s * -100%%)",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
				v.Value(),
			)
		}

		return &Properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateZPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s %s 1px",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}

func NegTranslateZPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(translatePropp): fmt.Sprintf("%s %s -1px",
				s.theme.UseVarKey(themeCustom, "translate-x"),
				s.theme.UseVarKey(themeCustom, "translate-y"),
			),
		}
	}
}
