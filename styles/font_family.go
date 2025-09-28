package styles

import (
	"fmt"

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

func FontFamilyCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func FontFamilyValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontFamilyProp): value.String(),
		}
	}
}
