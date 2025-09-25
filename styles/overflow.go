package styles

func OverflowAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp): "auto",
		}
	}
}

func OverflowHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp): "hidden",
		}
	}
}

func OverflowClip() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp): "clip",
		}
	}
}

func OverflowVisible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp): "visible",
		}
	}
}

func OverflowScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp): "scroll",
		}
	}
}

func OverflowXAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowXProp): "auto",
		}
	}
}

func OverflowYAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowYProp): "auto",
		}
	}
}

func OverflowXHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowXProp): "hidden",
		}
	}
}

func OverflowYHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowYProp): "hidden",
		}
	}
}

func OverflowXClip() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowXProp): "clip",
		}
	}
}

func OverflowYClip() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowYProp): "clip",
		}
	}
}

func OverflowXVisible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowXProp): "visible",
		}
	}
}

func OverflowYVisible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowYProp): "visible",
		}
	}
}

func OverflowXScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowXProp): "scroll",
		}
	}
}

func OverflowYScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowYProp): "scroll",
		}
	}
}
