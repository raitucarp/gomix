package styles

import (
	"github.com/raitucarp/gomix/value"
)

func Aspect(ratio value.Value) ApplyProp {
	return func(s *Style) StyleProp {

		return &Properties{
			string(aspectRatioProp): ratio.Value(),
		}
	}
}

func AspectSquare() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(aspectRatioProp): "1 / 1",
		}
	}
}

func AspectVideo() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(aspectRatioProp): s.theme.UseVarKey(themeAspect, "video"),
		}
	}
}

func AspectAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(aspectRatioProp): "auto",
		}
	}
}
