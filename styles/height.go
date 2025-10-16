package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func H(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func HFraction(fraction float32) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func HAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "auto",
		}
	}
}

func HPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "1px",
		}
	}
}

func HFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100%",
		}
	}
}

func HScreen() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100vh",
		}
	}
}

func HDvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100dvw",
		}
	}
}

func HDvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100dvh",
		}
	}
}

func HLvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100lvw",
		}
	}
}

func HLvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100lvh",
		}
	}
}

func HSvw() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100svw",
		}
	}
}

func HSvh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "100svh",
		}
	}
}

func HMin() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "min-content",
		}
	}
}

func HMax() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "max-content",
		}
	}
}

func HFit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "fit-content",
		}
	}
}

func HLh() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): "1lh",
		}
	}
}

func HValueBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(heightProp): val.Value(),
		}
	}
}
