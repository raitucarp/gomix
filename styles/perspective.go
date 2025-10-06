package styles

import (
	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func PerspectiveDramatic() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(theme.Perspective, "dramatic"),
		}
	}
}

func PerspectiveNear() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(theme.Perspective, "near"),
		}
	}
}

func PerspectiveNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(theme.Perspective, "normal"),
		}
	}
}

func PerspectiveMidrange() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(theme.Perspective, "midrange"),
		}
	}
}

func PerspectiveDistant() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(theme.Perspective, "distant"),
		}
	}
}

func PerspectiveNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): "none",
		}
	}
}

func Perspective(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): val.Value(),
		}
	}
}
