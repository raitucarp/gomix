package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Basis(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): val.Value(),
		}
	}
}

func BasisFraction(fraction float64) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func BasisFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): "100%",
		}
	}
}

func BasisAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): "auto",
		}
	}
}

func Basis3xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "3xs"),
		}
	}
}

func Basis2xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "2xs"),
		}
	}
}

func BasisXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "xs"),
		}
	}
}

func BasisSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "sm"),
		}
	}
}

func BasisMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "md"),
		}
	}
}

func BasisLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "lg"),
		}
	}
}

func BasisXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "xl"),
		}
	}
}

func Basis2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "2xl"),
		}
	}
}

func Basis3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "3xl"),
		}
	}
}

func Basis4xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "4xl"),
		}
	}
}

func Basis5xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "5xl"),
		}
	}
}

func Basis6xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "6xl"),
		}
	}

}

func Basis7xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexBasisProp): s.theme.UseVarKey(themeContainer, "7xl"),
		}
	}
}
