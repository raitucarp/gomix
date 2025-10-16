package styles

import "github.com/raitucarp/gomix/value"

func BgTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "top left",
		}
	}
}

func BgTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "top",
		}
	}
}

func BgTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "top right",
		}
	}
}

func BgLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "left",
		}
	}
}

func BgCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "center",
		}
	}
}

func BgRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "right",
		}
	}
}

func BgBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "bottom left",
		}
	}
}

func BgBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "bottom",
		}
	}
}

func BgBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): "bottom right",
		}
	}
}

func BgPosition(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundPositionProp): val.Value(),
		}
	}
}
