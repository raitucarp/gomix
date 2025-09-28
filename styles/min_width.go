package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func MinW(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func MinWFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func MinW3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "3xs"),
		}
	}
}

func MinW2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "2xs"),
		}
	}
}

func MinWXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "xs"),
		}
	}
}

func MinWSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "sm"),
		}
	}
}

func MinWMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "md"),
		}
	}
}

func MinWLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "lg"),
		}
	}
}

func MinWXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "xl"),
		}
	}
}

func MinW2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "2xl"),
		}
	}
}

func MinW3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "3xl"),
		}
	}
}

func MinW4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "4xl"),
		}
	}
}

func MinW5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "5xl"),
		}
	}
}

func MinW6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "6xl"),
		}
	}
}

func MinW7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(themes.Container, "7xl"),
		}
	}
}

func MinWAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "auto",
		}
	}
}

func MinWPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "1px",
		}
	}
}

func MinWFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100%",
		}
	}
}

func MinWScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100vw",
		}
	}
}

func MinWDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100dvw",
		}
	}
}

func MinWDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100dvh",
		}
	}
}

func MinWLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100lvw",
		}
	}
}

func MinWLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100lvh",
		}
	}
}

func MinWSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100svw",
		}
	}
}

func MinWSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "100svh",
		}
	}
}

func MinWMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "min-content",
		}
	}
}

func MinWMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "max-content",
		}
	}
}

func MinWFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): "fit-content",
		}
	}
}

func MinWCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func MinWValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): value.String(),
		}
	}
}
