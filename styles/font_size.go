package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func TextXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "xs"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "xs--line-height"),
		}
	}
}

func TextSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "sm"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "sm--line-height"),
		}
	}
}

func TextBase() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "base"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "base--line-height"),
		}
	}
}

func TextLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "lg"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "lg--line-height"),
		}
	}
}

func TextXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "xl--line-height"),
		}
	}
}

func Text2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "2xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "2xl--line-height"),
		}
	}
}

func Text3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "3xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "3xl--line-height"),
		}
	}
}

func Text4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "4xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "4xl--line-height"),
		}
	}
}

func Text5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "5xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "5xl--line-height"),
		}
	}
}

func Text6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "6xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "6xl--line-height"),
		}
	}
}

func Text7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "7xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "7xl--line-height"),
		}
	}
}

func Text8xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "8xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "8xl--line-height"),
		}
	}
}

func Text9xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp):   s.theme.UseVarKey(themes.Text, "9xl"),
			string(lineHeightProp): s.theme.UseVarKey(themes.Text, "9xl--line-height"),
		}
	}
}

func TextCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func TextValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp): value.String(),
		}
	}
}
