package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func ScaleNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scaleProp): "none",
		}
	}
}

func Scale(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%d%% * -1) calc(%d%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%d%% %d%%", v, v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#v%% * -1) calc(%#v%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%#v%% %#v%%", v, v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#v%% * -1) calc(%#v%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%#v%% %#v%%", v, v)
			}
		case value.CustomProperty:
			propValue = fmt.Sprintf("%s %s", v.Value(), v.Value())
		case value.LiteralValue:
			propValue = v.Value()
		case value.Value:
			propValue = v.Value()
		}

		return &Properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleX(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%d%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%d%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#v%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%#v%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%#v%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%#v%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case value.Value:
			propValue = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(theme.Custom, "scale-y"))
		}

		return &Properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleY(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%d%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %d%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%#v%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %#v%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%#v%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %#v%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case value.Value:
			propValue = fmt.Sprintf("%s %s", s.theme.UseVarKey(theme.Custom, "scale-x"), v.Value())
		}

		return &Properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleZ(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("%s %s calc(%d%% * -1)",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			} else {
				propValue = fmt.Sprintf("%s %s %d%%",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("%s %s calc(%#v%% * -1)",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			} else {
				propValue = fmt.Sprintf("%s %s %#v%%",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("%s %s calc(%#v%% * -1)",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			} else {
				propValue = fmt.Sprintf("%s %s %#v%%",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			}
		case value.Value:
			propValue = fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(theme.Custom, "scale-x"),
				s.theme.UseVarKey(theme.Custom, "scale-y"),
				v.Value(),
			)
		}

		return &Properties{
			string(scaleProp): propValue,
		}
	}
}

func Scale3d() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scaleProp): fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(theme.Custom, "scale-x"),
				s.theme.UseVarKey(theme.Custom, "scale-y"),
				s.theme.UseVarKey(theme.Custom, "scale-z"),
			),
		}
	}
}
