package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func W(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func WFraction(fraction float32) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func W3xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "3xs"),
		}
	}
}

func W2xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "2xs"),
		}
	}
}

func WXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "xs"),
		}
	}
}

func WSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "sm"),
		}
	}
}

func WMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "md"),
		}
	}
}

func WLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "lg"),
		}
	}
}

func WXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "xl"),
		}
	}
}

func W2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "2xl"),
		}
	}
}

func W3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "3xl"),
		}
	}
}

func W4xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "4xl"),
		}
	}
}

func W5xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "5xl"),
		}
	}
}

func W6xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "6xl"),
		}
	}
}

func W7xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): s.theme.UseVarKey(themeContainer, "7xl"),
		}
	}
}

func WAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "auto",
		}
	}
}

func WPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "1px",
		}
	}
}

func WFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100%",
		}
	}
}

func WScreen() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100vw",
		}
	}
}

func WDvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100dvw",
		}
	}
}

func WDvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100dvh",
		}
	}
}

func WLvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100lvw",
		}
	}
}

func WLvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100lvh",
		}
	}
}

func WSvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100svw",
		}
	}
}

func WSvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "100svh",
		}
	}
}

func WMin() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "min-content",
		}
	}
}

func WMax() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "max-content",
		}
	}
}

func WFit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): "fit-content",
		}
	}
}

func WBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp): val.Value(),
		}
	}
}

func Size(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  fmt.Sprintf("calc(var(--spacing) * %d)", number),
			string(heightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func SizeFraction(fraction float32) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  fmt.Sprintf("calc(var(--spacing) * %#v)", fraction),
			string(heightProp): fmt.Sprintf("calc(var(--spacing) * %#v)", fraction),
		}
	}
}

func SizeAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "auto",
			string(heightProp): "auto",
		}
	}
}

func SizePx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "1px",
			string(heightProp): "1px",
		}
	}
}

func SizeDvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100dvw",
			string(heightProp): "100dvw",
		}
	}
}

func SizeDvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100dvh",
			string(heightProp): "100dvh",
		}
	}
}

func SizeLvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100lvw",
			string(heightProp): "100lvw",
		}
	}
}

func SizeLvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100lvh",
			string(heightProp): "100lvh",
		}
	}
}

func SizeSvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100svw",
			string(heightProp): "100svw",
		}
	}
}

func SizeSvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "100svh",
			string(heightProp): "100svh",
		}
	}
}

func SizeMin() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "min-content",
			string(heightProp): "min-content",
		}
	}
}

func SizeMax() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "max-content",
			string(heightProp): "max-content",
		}
	}
}

func SizeFit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  "fit-content",
			string(heightProp): "fit-content",
		}
	}
}

func SizeBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(widthProp):  val.Value(),
			string(heightProp): val.Value(),
		}
	}
}
