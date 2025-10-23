package styles

import (
	"github.com/raitucarp/gomix/value"
)

func EaseLinear() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): "linear",
		}
	}
}

func EaseIn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): s.theme.UseVarKey(themeEase, "in"),
		}
	}
}

func EaseOut() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): s.theme.UseVarKey(themeEase, "out"),
		}
	}
}

func EaseInOut() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): s.theme.UseVarKey(themeEase, "in-out"),
		}
	}
}

func EaseInitial() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): s.theme.UseVarKey(themeEase, "initial"),
		}
	}
}

func Ease(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transitionTimingProp): val.Value(),
		}
	}
}
