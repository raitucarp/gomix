package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func addLineHeight(s *style, prop *properties, size string, lineHeights ...any) {
	if len(lineHeights) <= 0 {
		(*prop)[string(lineHeightProp)] = s.theme.UseVarKey(theme.Text, size+"--line-height")
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
	return func(s *style) styleProp {
		size := "xs"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextSm(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "sm"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)
		return prop
	}
}

func TextBase(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "base"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextLg(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "lg"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextXl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text2xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "2xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text3xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "3xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text4xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "4xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text5xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "5xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text6xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "6xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text7xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "7xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text8xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "8xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func Text9xl(lineHeights ...any) ApplyProp {
	return func(s *style) styleProp {
		size := "9xl"
		prop := &properties{
			string(fontSizeProp): s.theme.UseVarKey(theme.Text, size),
		}

		addLineHeight(s, prop, size, lineHeights...)

		return prop
	}
}

func TextSizeBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontSizeProp): val.Value(),
		}
	}
}
