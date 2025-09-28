package styles

import "fmt"

func FontThin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "100",
		}
	}
}

func FontExtraLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "200",
		}
	}
}

func FontLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "300",
		}
	}
}

func FontNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "400",
		}
	}
}

func FontMedium() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "500",
		}
	}
}

func FontSemibold() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "600",
		}
	}
}

func FontBold() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "700",
		}
	}
}

func FontExtraBold() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "800",
		}
	}
}

func FontBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): "900",
		}
	}
}

func FontWeightCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func FontWeightValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontWeightProp): value.String(),
		}
	}
}
