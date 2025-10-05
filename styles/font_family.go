package styles

import "github.com/raitucarp/gomix/theme"

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

func FontFamilyBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): value.Value(),
		}
	}
}
