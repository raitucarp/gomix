package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Transform(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformProp): val.Value(),
		}
	}
}

func TransformNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformProp): "none",
		}
	}
}

func TransformGpu() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformProp): fmt.Sprintf("translateZ(0) %s %s %s %s %s",
				s.theme.UseVarKey(themeCustom, "rotate-x"),
				s.theme.UseVarKey(themeCustom, "rotate-y"),
				s.theme.UseVarKey(themeCustom, "rotate-z"),
				s.theme.UseVarKey(themeCustom, "skew-x"),
				s.theme.UseVarKey(themeCustom, "skew-y"),
			),
		}
	}
}

func TransformCpu() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformProp): fmt.Sprintf("%s %s %s %s %s",
				s.theme.UseVarKey(themeCustom, "rotate-x"),
				s.theme.UseVarKey(themeCustom, "rotate-y"),
				s.theme.UseVarKey(themeCustom, "rotate-z"),
				s.theme.UseVarKey(themeCustom, "skew-x"),
				s.theme.UseVarKey(themeCustom, "skew-y"),
			),
		}
	}
}
