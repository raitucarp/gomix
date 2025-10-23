package styles

import (
	"github.com/raitucarp/gomix/value"
)

func PerspectiveDramatic() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): s.theme.UseVarKey(themePerspective, "dramatic"),
		}
	}
}

func PerspectiveNear() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): s.theme.UseVarKey(themePerspective, "near"),
		}
	}
}

func PerspectiveNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): s.theme.UseVarKey(themePerspective, "normal"),
		}
	}
}

func PerspectiveMidrange() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): s.theme.UseVarKey(themePerspective, "midrange"),
		}
	}
}

func PerspectiveDistant() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): s.theme.UseVarKey(themePerspective, "distant"),
		}
	}
}

func PerspectiveNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): "none",
		}
	}
}

func Perspective(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveProp): val.Value(),
		}
	}
}
