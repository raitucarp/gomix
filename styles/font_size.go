package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func addLineHeight(s *Style, prop *Properties, size string, lineHeights ...any) {
	if len(lineHeights) <= 0 {
		(*prop)[string(lineHeightProp)] = s.theme.UseVarKey(themeText, size+"--line-height")
		return
	}

	lineHeight := lineHeights[0]

	switch v := lineHeight.(type) {
	case int:
		(*prop)[string(lineHeightProp)] = fmt.Sprintf("calc(var(--spacing) * %d)", v)
	case string:
		(*prop)[string(lineHeightProp)] = fmt.Sprintf("var(--%s)", v)
	case value.Value:
		(*prop)[string(lineHeightProp)] = fmt.Sprintf("var(--%s)", v)
	}

}

func TextXs(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "xs"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextSm(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "sm"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)
		return prop
	}
}

func TextBase(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "base"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextLg(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "lg"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextXl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text2xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "2xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text3xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "3xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text4xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "4xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text5xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "5xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text6xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "6xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text7xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "7xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text8xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "8xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text9xl(lineHeights ...any) ApplyProp {
	return func(s *Style) StyleProp {
		size := "9xl"
		prop := &Properties{
			string(fontSizeProp): s.theme.UseVarKey(themeText, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextSizeBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontSizeProp): val.Value(),
		}
	}
}
