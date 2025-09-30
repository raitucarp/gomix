package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func W(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func WFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func W3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "3xs"),
		}
	}
}

func W2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "2xs"),
		}
	}
}

func WXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "xs"),
		}
	}
}

func WSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "sm"),
		}
	}
}

func WMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "md"),
		}
	}
}

func WLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "lg"),
		}
	}
}

func WXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "xl"),
		}
	}
}

func W2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "2xl"),
		}
	}
}

func W3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "3xl"),
		}
	}
}

func W4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "4xl"),
		}
	}
}

func W5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "5xl"),
		}
	}
}

func W6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "6xl"),
		}
	}
}

func W7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): s.theme.UseVarKey(themes.Container, "7xl"),
		}
	}
}

func WAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "auto",
		}
	}
}

func WPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "1px",
		}
	}
}

func WFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100%",
		}
	}
}

func WScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100vw",
		}
	}
}

func WDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100dvw",
		}
	}
}

func WDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100dvh",
		}
	}
}

func WLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100lvw",
		}
	}
}

func WLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100lvh",
		}
	}
}

func WSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100svw",
		}
	}
}

func WSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "100svh",
		}
	}
}

func WMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "min-content",
		}
	}
}

func WMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "max-content",
		}
	}
}

func WFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): "fit-content",
		}
	}
}

func WBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp): value.Value(),
		}
	}
}

func Size(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  fmt.Sprintf("calc(var(--spacing) * %d)", number),
			string(heightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func SizeFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  fmt.Sprintf("calc(var(--spacing) * %.2f)", fraction),
			string(heightProp): fmt.Sprintf("calc(var(--spacing) * %.2f)", fraction),
		}
	}
}

func SizeAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "auto",
			string(heightProp): "auto",
		}
	}
}

func SizePx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "1px",
			string(heightProp): "1px",
		}
	}
}

func SizeDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100dvw",
			string(heightProp): "100dvw",
		}
	}
}

func SizeDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100dvh",
			string(heightProp): "100dvh",
		}
	}
}

func SizeLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100lvw",
			string(heightProp): "100lvw",
		}
	}
}

func SizeLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100lvh",
			string(heightProp): "100lvh",
		}
	}
}

func SizeSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100svw",
			string(heightProp): "100svw",
		}
	}
}

func SizeSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "100svh",
			string(heightProp): "100svh",
		}
	}
}

func SizeMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "min-content",
			string(heightProp): "min-content",
		}
	}
}

func SizeMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "max-content",
			string(heightProp): "max-content",
		}
	}
}

func SizeFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  "fit-content",
			string(heightProp): "fit-content",
		}
	}
}

func SizeBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(widthProp):  value.Value(),
			string(heightProp): value.Value(),
		}
	}
}
