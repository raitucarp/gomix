package styles

import "fmt"

func Gap(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): fmt.Sprintf("calc(var(--spacing)) * %d", number),
		}
	}
}

func GapBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapProp): value.Value(),
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

func GapXBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapXProp): value.Value(),
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

func GapYBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gapYProp): value.Value(),
		}
	}
}
