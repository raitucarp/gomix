package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Transition(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		if len(value) == 1 {
			return &properties{
				string(transitionPropertyProp): value[0].Value(),
				string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
				string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
			}
		}

		return &properties{
			string(transitionPropertyProp): fmt.Sprintf("color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, %s, %s, %s, opacity, box-shadow, transform, translate, scale, rotate, filter, -webkit-backdrop-filter, backdrop-filter, display, content-visibility, overlay, pointer-events",
				s.theme.UseVarKey(theme.Custom, "gradient-from"),
				s.theme.UseVarKey(theme.Custom, "gradient-via"),
				s.theme.UseVarKey(theme.Custom, "gradient-to"),
			),
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionAll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): "all",
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionColors() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): fmt.Sprintf("color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, %s, %s, %s", s.theme.UseVarKey(theme.Custom, "gradient-from"), s.theme.UseVarKey(theme.Custom, "gradient-via"), s.theme.UseVarKey(theme.Custom, "gradient-to")),
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionOpacity() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): "opacity",
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionShadow() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): "box-shadow",
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionTransform() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): "transform, translate, scale, rotate",
			string(transitionTimingProp):   s.theme.UseVarKey(theme.Custom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(theme.Custom, "default-transition-duration"),
		}
	}
}

func TransitionNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionPropertyProp): "none",
		}
	}
}
