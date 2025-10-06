package styles

import (
	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func FontSans() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(theme.Font, "sans"),
		}
	}
}

func FontSerif() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(theme.Font, "serif"),
		}
	}
}

func FontMono() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(theme.Font, "mono"),
		}
	}
}

func FontFamilyBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): val.Value(),
		}
	}
}
