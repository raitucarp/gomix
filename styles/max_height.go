package styles

import (
	"fmt"
)

func MaxH(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func MaxHFraction(fraction float32) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): fmt.Sprintf("calc(%.2f * 100%%)", fraction),
		}
	}
}

func MaxHNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "none",
		}
	}
}

func MaxHPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "1px",
		}
	}
}

func MaxHFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100%",
		}
	}
}

func MaxHScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100vh",
		}
	}
}

func MaxHDvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100dvw",
		}
	}
}

func MaxHDvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100dvh",
		}
	}
}

func MaxHLvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100lvw",
		}
	}
}

func MaxHLvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100lvh",
		}
	}
}

func MaxHSvw() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100svw",
		}
	}
}

func MaxHSvh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "100svh",
		}
	}
}

func MaxHMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "min-content",
		}
	}
}

func MaxHMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "max-content",
		}
	}
}

func MaxHFit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "fit-content",
		}
	}
}

func MaxHLh() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): "1lh",
		}
	}
}

func MaxHBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maxHeightProp): value.Value(),
		}
	}
}
