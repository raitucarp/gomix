package styles

import (
	"fmt"
)

func SnapNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapTypeProp): "none",
		}
	}
}

func SnapX() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapTypeProp): fmt.Sprintf("x %s", s.theme.UseVarKey(themeCustom, "scroll-snap-strictness")),
		}
	}
}

func SnapY() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapTypeProp): fmt.Sprintf("y %s", s.theme.UseVarKey(themeCustom, "scroll-snap-strictness")),
		}
	}
}

func SnapBoth() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapTypeProp): fmt.Sprintf("both %s", s.theme.UseVarKey(themeCustom, "scroll-snap-strictness")),
		}
	}
}

func SnapMandatory() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapStrictnessProp): "mandatory",
		}
	}
}

func SnapProximity() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapStrictnessProp): "proximity",
		}
	}
}
