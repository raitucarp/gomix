package styles

import "github.com/raitucarp/gomix/themes"

func PerspectiveDramatic() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(themes.Perspective, "dramatic"),
		}
	}
}

func PerspectiveNear() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(themes.Perspective, "near"),
		}
	}
}

func PerspectiveNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(themes.Perspective, "normal"),
		}
	}
}

func PerspectiveMidrange() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(themes.Perspective, "midrange"),
		}
	}
}

func PerspectiveDistant() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): s.theme.UseVarKey(themes.Perspective, "distant"),
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

func Perspective(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveProp): value.Value(),
		}
	}
}
