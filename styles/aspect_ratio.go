package styles

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/raitucarp/gomix/themes"
)

func Aspect(ratio string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(aspectRatioProp): ratio,
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
			string(aspectRatioProp): s.theme.UseVarKey(themes.Aspect, "video"),
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

func AspectCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(aspectRatioProp): fmt.Sprintf("var(--%s)", strcase.ToDelimited(customProperty, '-')),
		}
	}
}
