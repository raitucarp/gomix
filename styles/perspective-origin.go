package styles

import "github.com/raitucarp/gomix/value"

func PerspectiveOriginCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "center",
		}
	}
}

func PerspectiveOriginTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "top",
		}
	}
}

func PerspectiveOriginTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "top right",
		}
	}
}

func PerspectiveOriginRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "right",
		}
	}
}

func PerspectiveOriginBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "bottom right",
		}
	}
}

func PerspectiveOriginBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "bottom",
		}
	}
}

func PerspectiveOriginBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "bottom left",
		}
	}
}

func PerspectiveOriginLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "left",
		}
	}
}

func PerspectiveOriginTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): "top left",
		}
	}
}

func PerspectiveOrigin(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(perspectiveOriginProp): val.Value(),
		}
	}
}
