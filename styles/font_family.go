package styles

import (
	"github.com/raitucarp/gomix/themes"
)

func FontSans() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(themes.Font, "sans"),
		}
	}
}

func FontSerif() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(themes.Font, "serif"),
		}
	}
}

func FontMono() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): s.theme.UseVarKey(themes.Font, "mono"),
		}
	}
}

func FontFamilyBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): value.Value(),
		}
	}
}
