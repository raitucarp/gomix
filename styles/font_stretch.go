package styles

import "fmt"

func FontStretchUltraCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "ultra-condensed",
		}
	}
}

func FontStretchExtraCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "extra-condensed",
		}
	}
}

func FontStretchCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "condensed",
		}
	}
}

func FontStretchSemiCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "semi-condensed",
		}
	}
}

func FontStretchNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "normal",
		}
	}
}

func FontStretchSemiExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "semi-expanded",
		}
	}
}

func FontStretchExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "expanded",
		}
	}
}

func FontStretchExtraExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "extra-expanded",
		}
	}
}

func FontStretchUltraExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "ultra-expanded",
		}
	}
}

func FontStretchPercentage(percentage int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): fmt.Sprintf("%d%%", percentage),
		}
	}
}

func FontStretchCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func FontStretchValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): value.String(),
		}
	}
}
