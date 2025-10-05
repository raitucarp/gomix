package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func MaxW(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func MaxWFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func MaxW3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "3xs"),
		}
	}
}

func MaxW2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "2xs"),
		}
	}
}

func MaxWXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "xs"),
		}
	}
}

func MaxWSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "sm"),
		}
	}
}

func MaxWMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "md"),
		}
	}
}

func MaxWLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "lg"),
		}
	}
}

func MaxWXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "xl"),
		}
	}
}

func MaxW2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "2xl"),
		}
	}
}

func MaxW3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "3xl"),
		}
	}
}

func MaxW4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "4xl"),
		}
	}
}

func MaxW5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "5xl"),
		}
	}
}

func MaxW6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "6xl"),
		}
	}
}

func MaxW7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): s.theme.UseVarKey(theme.Container, "7xl"),
		}
	}
}

func MaxWNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "none",
		}
	}
}

func MaxWPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "1px",
		}
	}
}

func MaxWFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100%",
		}
	}
}

func MaxWScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100vw",
		}
	}
}

func MaxWDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100dvw",
		}
	}
}

func MaxWDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100dvh",
		}
	}
}

func MaxWLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100lvw",
		}
	}
}

func MaxWLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100lvh",
		}
	}
}

func MaxWSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100svw",
		}
	}
}

func MaxWSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "100svh",
		}
	}
}

func MaxWMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "min-content",
		}
	}
}

func MaxWMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "max-content",
		}
	}
}

func MaxWFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): "fit-content",
		}
	}
}

func Container() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100%",
			string(mediaQueryProp(maxWidthProp, "width >= 40rem").String()): "40rem",
			string(mediaQueryProp(maxWidthProp, "width >= 48rem").String()): "48rem",
			string(mediaQueryProp(maxWidthProp, "width >= 64rem").String()): "64rem",
			string(mediaQueryProp(maxWidthProp, "width >= 80rem").String()): "80rem",
			string(mediaQueryProp(maxWidthProp, "width >= 96rem").String()): "96rem",
		}
	}
}

func MaxWBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxWidthProp): value.Value(),
		}
	}
}
