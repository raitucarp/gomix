package styles

import "github.com/raitucarp/gomix/value"

func OriginCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "center",
		}
	}
}

func OriginTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "top",
		}
	}
}

func OriginTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "top right",
		}
	}
}

func OriginRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "right",
		}
	}
}

func OriginBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "bottom right",
		}
	}
}

func OriginBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "bottom",
		}
	}
}

func OriginBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "bottom left",
		}
	}
}

func OriginLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "left",
		}
	}
}

func OriginTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): "top left",
		}
	}
}

func Origin(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformOriginProp): val.Value(),
		}
	}
}
