package styles

func OverflowAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp): "auto",
		}
	}
}

func OverflowHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp): "hidden",
		}
	}
}

func OverflowClip() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp): "clip",
		}
	}
}

func OverflowVisible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp): "visible",
		}
	}
}

func OverflowScroll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp): "scroll",
		}
	}
}

func OverflowXAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowXProp): "auto",
		}
	}
}

func OverflowYAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowYProp): "auto",
		}
	}
}

func OverflowXHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowXProp): "hidden",
		}
	}
}

func OverflowYHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowYProp): "hidden",
		}
	}
}

func OverflowXClip() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowXProp): "clip",
		}
	}
}

func OverflowYClip() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowYProp): "clip",
		}
	}
}

func OverflowXVisible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowXProp): "visible",
		}
	}
}

func OverflowYVisible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowYProp): "visible",
		}
	}
}

func OverflowXScroll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowXProp): "scroll",
		}
	}
}

func OverflowYScroll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowYProp): "scroll",
		}
	}
}
