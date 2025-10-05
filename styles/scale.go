package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func ScaleNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scaleProp): "none",
		}
	}
}

func Scale(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%d%% * -1) calc(%d%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%d%% %d%%", v, v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2f%% * -1) calc(%.2f%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%.2f%% %.2f%%", v, v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2f%% * -1) calc(%.2f%% * -1)", v, v)
			} else {
				propValue = fmt.Sprintf("%.2f%% %.2f%%", v, v)
			}
		case customVariableProp:
			propValue = fmt.Sprintf("%s %s", v.Value(), v.Value())
		case val:
			propValue = v.Value()
		case customValue:
			propValue = v.Value()
		}

		return &properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleX(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%d%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%d%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2f%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%.2f%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("calc(%.2f%% * -1) %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			} else {
				propValue = fmt.Sprintf("%.2f%% %s", v, s.theme.UseVarKey(theme.Custom, "scale-y"))
			}
		case customValue:
			propValue = fmt.Sprintf("%s %s", v.Value(), s.theme.UseVarKey(theme.Custom, "scale-y"))
		}

		return &properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleY(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%d%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %d%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case float32:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%.2f%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %.2f%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("%s calc(%.2f%% * -1)", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			} else {
				propValue = fmt.Sprintf("%s %.2f%%", s.theme.UseVarKey(theme.Custom, "scale-x"), v)
			}
		case customValue:
			propValue = fmt.Sprintf("%s %s", s.theme.UseVarKey(theme.Custom, "scale-x"), v.Value())
		}

		return &properties{
			string(scaleProp): propValue,
		}
	}
}

func ScaleZ(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
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
				propValue = fmt.Sprintf("%s %s calc(%.2f%% * -1)",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			} else {
				propValue = fmt.Sprintf("%s %s %.2f%%",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			}
		case float64:
			if v < 0 {
				propValue = fmt.Sprintf("%s %s calc(%.2f%% * -1)",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			} else {
				propValue = fmt.Sprintf("%s %s %.2f%%",
					s.theme.UseVarKey(theme.Custom, "scale-x"),
					s.theme.UseVarKey(theme.Custom, "scale-y"),
					v,
				)
			}
		case customValue:
			propValue = fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(theme.Custom, "scale-x"),
				s.theme.UseVarKey(theme.Custom, "scale-y"),
				v.Value(),
			)
		}

		return &properties{
			string(scaleProp): propValue,
		}
	}
}

func Scale3d() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scaleProp): fmt.Sprintf("%s %s %s",
				s.theme.UseVarKey(theme.Custom, "scale-x"),
				s.theme.UseVarKey(theme.Custom, "scale-y"),
				s.theme.UseVarKey(theme.Custom, "scale-z"),
			),
		}
	}
}
