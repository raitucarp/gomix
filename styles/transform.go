package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Transform(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformProp): value.Value(),
		}
	}
}

func TransformNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformProp): "none",
		}
	}
}

func TransformGpu() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformProp): fmt.Sprintf("translateZ(0) %s %s %s %s %s",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				s.theme.UseVarKey(theme.Custom, "rotate-z"),
				s.theme.UseVarKey(theme.Custom, "skew-x"),
				s.theme.UseVarKey(theme.Custom, "skew-y"),
			),
		}
	}
}

func TransformCpu() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformProp): fmt.Sprintf("%s %s %s %s %s",
				s.theme.UseVarKey(theme.Custom, "rotate-x"),
				s.theme.UseVarKey(theme.Custom, "rotate-y"),
				s.theme.UseVarKey(theme.Custom, "rotate-z"),
				s.theme.UseVarKey(theme.Custom, "skew-x"),
				s.theme.UseVarKey(theme.Custom, "skew-y"),
			),
		}
	}
}
