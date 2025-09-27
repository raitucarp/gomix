package styles

import "fmt"

func Gap(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func GapValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): value.String(),
		}
	}
}

func GapX(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapXCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func GapXValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): value.String(),
		}
	}
}

func GapY(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapYCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func GapYValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): value.String(),
		}
	}
}
