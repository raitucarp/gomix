package styles

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/raitucarp/gomix/themes"
)

func (s *Styles) Aspect(ratio string) *Styles {
	s.addProp(AspectRatio, ratio)
	return s
}

func (s *Styles) AspectSquare() *Styles {
	s.addProp(AspectRatio, "1 / 1")
	return s
}

func (s *Styles) AspectVideo() *Styles {
	s.addProp(AspectRatio, s.theme.UseVarKey(themes.Aspect, "video"))
	return s
}

func (s *Styles) AspectAuto() *Styles {
	s.addProp(AspectRatio, "auto")
	return s
}

func (s *Styles) AspectProperty(varname string) *Styles {
	s.addProp(AspectRatio,
		fmt.Sprintf("var(--%s)", strcase.ToDelimited(varname, '-')),
	)
	return s
}
