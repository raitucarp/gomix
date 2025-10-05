package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Basis(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): value.Value(),
		}
	}
}

func BasisFraction(fraction float64) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func BasisFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): "100%",
		}
	}
}

func BasisAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): "auto",
		}
	}
}

func Basis3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "3xs"),
		}
	}
}

func Basis2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "2xs"),
		}
	}
}

func BasisXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "xs"),
		}
	}
}

func BasisSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "sm"),
		}
	}
}

func BasisMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "md"),
		}
	}
}

func BasisLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "lg"),
		}
	}
}

func BasisXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "xl"),
		}
	}
}

func Basis2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "2xl"),
		}
	}
}

func Basis3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "3xl"),
		}
	}
}

func Basis4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "4xl"),
		}
	}
}

func Basis5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "5xl"),
		}
	}
}

func Basis6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "6xl"),
		}
	}

}

func Basis7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexBasisProp): s.theme.UseVarKey(theme.Container, "7xl"),
		}
	}
}
