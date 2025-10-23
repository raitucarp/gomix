package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Transition(val ...value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		if len(val) == 1 {
			return &Properties{
				string(transitionPropertyProp): val[0].Value(),
				string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
				string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
			}
		}

		return &Properties{
			string(transitionPropertyProp): fmt.Sprintf("color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, %s, %s, %s, opacity, box-shadow, transform, translate, scale, rotate, filter, -webkit-backdrop-filter, backdrop-filter, display, content-visibility, overlay, pointer-events",
				s.theme.UseVarKey(themeCustom, "gradient-from"),
				s.theme.UseVarKey(themeCustom, "gradient-via"),
				s.theme.UseVarKey(themeCustom, "gradient-to"),
			),
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionAll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): "all",
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionColors() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): fmt.Sprintf("color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, %s, %s, %s", s.theme.UseVarKey(themeCustom, "gradient-from"), s.theme.UseVarKey(themeCustom, "gradient-via"), s.theme.UseVarKey(themeCustom, "gradient-to")),
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionOpacity() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): "opacity",
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionShadow() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): "box-shadow",
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionTransform() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): "transform, translate, scale, rotate",
			string(transitionTimingProp):   s.theme.UseVarKey(themeCustom, "default-transition-timing-function"),
			string(transitionDurationProp): s.theme.UseVarKey(themeCustom, "default-transition-duration"),
		}
	}
}

func TransitionNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionPropertyProp): "none",
		}
	}
}
