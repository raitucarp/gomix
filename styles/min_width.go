package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
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
			string(minWidthProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func MinW3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "3xs"),
		}
	}
}

func MinW2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "2xs"),
		}
	}
}

func MinWXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "xs"),
		}
	}
}

func MinWSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "sm"),
		}
	}
}

func MinWMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "md"),
		}
	}
}

func MinWLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "lg"),
		}
	}
}

func MinWXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "xl"),
		}
	}
}

func MinW2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "2xl"),
		}
	}
}

func MinW3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "3xl"),
		}
	}
}

func MinW4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "4xl"),
		}
	}
}

func MinW5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "5xl"),
		}
	}
}

func MinW6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "6xl"),
		}
	}
}

func MinW7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): s.theme.UseVarKey(theme.Container, "7xl"),
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

func MinWBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minWidthProp): val.Value(),
		}
	}
}
