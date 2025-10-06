package styles

import (
	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func Aspect(ratio value.Value) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(aspectRatioProp): ratio.Value(),
		}
	}
}

func AspectSquare() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(aspectRatioProp): "1 / 1",
		}
	}
}

func AspectVideo() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(aspectRatioProp): s.theme.UseVarKey(theme.Aspect, "video"),
		}
	}
}

func AspectAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(aspectRatioProp): "auto",
		}
	}
}
