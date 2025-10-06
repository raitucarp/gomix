package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func MinH(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func MinHFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): fmt.Sprintf("calc(%#v * 100%%)", fraction),
		}
	}
}

func MinHAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "auto",
		}
	}
}

func MinHPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "1px",
		}
	}
}

func MinHFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100%",
		}
	}
}

func MinHScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100vh",
		}
	}
}

func MinHDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100dvw",
		}
	}
}

func MinHDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100dvh",
		}
	}
}

func MinHLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100lvw",
		}
	}
}

func MinHLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100lvh",
		}
	}
}

func MinHSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100svw",
		}
	}
}

func MinHSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "100svh",
		}
	}
}

func MinHMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "min-content",
		}
	}
}

func MinHMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "max-content",
		}
	}
}

func MinHFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "fit-content",
		}
	}
}

func MinHLh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): "1lh",
		}
	}
}

func MinHBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(minHeightProp): val.Value(),
		}
	}
}
