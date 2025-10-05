package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func TranslateNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): "none",
		}
	}
}

func Translate(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * %d) calc(var(--spacing) * %d)", v, v)
		case ratio:
			propValue = fmt.Sprintf("calc(%s * 100%%) calc(%s * 100%%)", v.Value(), v.Value())
		case customValue:
			propValue = fmt.Sprintf("%s %s", v, v)
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslate(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * -%d) calc(var(--spacing) * -%d)", v, v)
		case ratio:
			propValue = fmt.Sprintf("calc(%s * -100%%) calc(%s * -100%%)", v.Value(), v.Value())
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): `100% 100%`,
		}
	}
}

func NegTranslateFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): `-100% -100%`,
		}
	}
}

func TranslatePx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): `1px 1px`,
		}
	}
}

func NegTranslatePx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): `-1px -1px`,
		}
	}
}

func TranslateX(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * %d) %s",
				v,
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			)
		case ratio:
			propValue = fmt.Sprintf("calc(%s * 100%%) %s", v.Value(), s.theme.UseVarKey(theme.Custom, "translate-y"))
		case customValue:
			propValue = fmt.Sprintf("%s %s", v, s.theme.UseVarKey(theme.Custom, "translate-y"))
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateX(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("calc(var(--spacing) * -%d) %s",
				v,
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			)
		case ratio:
			propValue = fmt.Sprintf("calc(%s * -100%%) %s", v.Value(), s.theme.UseVarKey(theme.Custom, "translate-y"))
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateXFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("100%% %s",
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}

func NegTranslateXFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("-100%% %s",
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}

func TranslateXPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("1px %s",
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}

func NegTranslateXPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("-1px %s",
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}

func TranslateY(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s calc(var(--spacing) * %d)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				v,
			)
		case ratio:
			propValue = fmt.Sprintf("%s calc(%s * 100%%)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				v.Value(),
			)
		case customValue:
			propValue = fmt.Sprintf("%s %s",
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v,
			)
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateY(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s calc(var(--spacing) * -%d)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				v,
			)
		case ratio:
			propValue = fmt.Sprintf("%s calc(%s * -100%%)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				v.Value(),
			)
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateYFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s 100%%",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
			),
		}
	}
}

func NegTranslateYFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s -100%%",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
			),
		}
	}
}

func TranslateYPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s 1px",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
			),
		}
	}
}

func NegTranslateYPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s -1px",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
			),
		}
	}
}

func TranslateZ(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s calc(var(--spacing) * %d)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v,
			)
		case ratio:
			propValue = fmt.Sprintf("%s %s calc(%s * 100%%)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v.Value(),
			)
		case customValue:
			propValue = fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v,
			)
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func NegTranslateZ(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("%s %s calc(var(--spacing) * -%d)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v,
			)
		case ratio:
			propValue = fmt.Sprintf("%s %s calc(%s * -100%%)",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
				v.Value(),
			)
		}

		return &properties{
			string(translatePropp): propValue,
		}
	}
}

func TranslateZPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s %s 1px",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}

func NegTranslateZPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(translatePropp): fmt.Sprintf("%s %s -1px",
				s.theme.UseVarKey(theme.Custom, "translate-x"),
				s.theme.UseVarKey(theme.Custom, "translate-y"),
			),
		}
	}
}
