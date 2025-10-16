package styles

import "github.com/raitucarp/gomix/value"

func ObjectTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "top left",
		}
	}
}

func ObjectTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "top",
		}
	}
}

func ObjectTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "top right",
		}
	}
}

func ObjectLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "left",
		}
	}
}

func ObjectCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "center",
		}
	}
}

func ObjectRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "right",
		}
	}
}

func ObjectBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "bottom left",
		}
	}
}

func ObjectBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "bottom",
		}
	}
}

func ObjectBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): "bottom right",
		}
	}
}

func ObjectBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectPositionProp): val.Value(),
		}
	}
}
