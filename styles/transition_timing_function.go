package styles

import "github.com/raitucarp/gomix/theme"

func EaseLinear() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): "linear",
		}
	}
}

func EaseIn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): s.theme.UseVarKey(theme.Ease, "in"),
		}
	}
}

func EaseOut() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): s.theme.UseVarKey(theme.Ease, "out"),
		}
	}
}

func EaseInOut() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): s.theme.UseVarKey(theme.Ease, "in-out"),
		}
	}
}

func EaseInitial() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): s.theme.UseVarKey(theme.Ease, "initial"),
		}
	}
}

func Ease(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionTimingProp): value.Value(),
		}
	}
}
