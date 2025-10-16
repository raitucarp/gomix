package styles

import "github.com/raitucarp/gomix/value"

func MaskTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "top left",
		}
	}
}

func MaskTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "top",
		}
	}

}

func MaskTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "top right",
		}
	}
}

func MaskLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "left",
		}
	}
}

func MaskCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "center",
		}
	}
}

func MaskRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "right",
		}
	}
}

func MaskBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "bottom left",
		}
	}
}

func MaskBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "bottom",
		}
	}
}

func MaskBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): "bottom right",
		}
	}
}

func MaskPosition(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskPositionProp): val.Value(),
		}
	}
}
