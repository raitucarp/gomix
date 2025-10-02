package styles

func PerspectiveOriginCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "center",
		}
	}
}

func PerspectiveOriginTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "top",
		}
	}
}

func PerspectiveOriginTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "top right",
		}
	}
}

func PerspectiveOriginRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "right",
		}
	}
}

func PerspectiveOriginBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "bottom right",
		}
	}
}

func PerspectiveOriginBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "bottom",
		}
	}
}

func PerspectiveOriginBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "bottom left",
		}
	}
}

func PerspectiveOriginLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "left",
		}
	}
}

func PerspectiveOriginTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): "top left",
		}
	}
}

func PerspectiveOrigin(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(perspectiveOriginProp): value.Value(),
		}
	}
}
