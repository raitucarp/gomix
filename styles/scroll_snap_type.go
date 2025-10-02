package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func SnapNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapTypeProp): "none",
		}
	}
}

func SnapX() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapTypeProp): fmt.Sprintf("x %s", s.theme.UseVarKey(themes.Custom, "scroll-snap-strictness")),
		}
	}
}

func SnapY() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapTypeProp): fmt.Sprintf("y %s", s.theme.UseVarKey(themes.Custom, "scroll-snap-strictness")),
		}
	}
}

func SnapBoth() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapTypeProp): fmt.Sprintf("both %s", s.theme.UseVarKey(themes.Custom, "scroll-snap-strictness")),
		}
	}
}

func SnapMandatory() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapStrictnessProp): "mandatory",
		}
	}
}

func SnapProximity() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapStrictnessProp): "proximity",
		}
	}
}
