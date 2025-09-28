package styles

import "fmt"

func BgTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "top left",
		}
	}
}

func BgTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "top",
		}
	}
}

func BgTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "top right",
		}
	}
}

func BgLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "left",
		}
	}
}

func BgCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "center",
		}
	}
}

func BgRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "right",
		}
	}
}

func BgBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "bottom left",
		}
	}
}

func BgBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "bottom",
		}
	}
}

func BgBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): "bottom right",
		}
	}
}

func BgPositionCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func BgPositionValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): value.String(),
		}
	}
}
