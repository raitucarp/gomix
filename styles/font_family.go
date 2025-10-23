package styles

import (
	"github.com/raitucarp/gomix/value"
)

func FontSans() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontFamilyProp): s.theme.UseVarKey(themeFont, "sans"),
		}
	}
}

func FontSerif() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontFamilyProp): s.theme.UseVarKey(themeFont, "serif"),
		}
	}
}

func FontMono() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontFamilyProp): s.theme.UseVarKey(themeFont, "mono"),
		}
	}
}

func FontFamilyBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontFamilyProp): val.Value(),
		}
	}
}
