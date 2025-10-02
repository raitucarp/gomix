package styles

func OriginCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "center",
		}
	}
}

func OriginTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "top",
		}
	}
}

func OriginTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "top right",
		}
	}
}

func OriginRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "right",
		}
	}
}

func OriginBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "bottom right",
		}
	}
}

func OriginBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "bottom",
		}
	}
}

func OriginBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "bottom left",
		}
	}
}

func OriginLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "left",
		}
	}
}

func OriginTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): "top left",
		}
	}
}

func Origin(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformOriginProp): value.Value(),
		}
	}
}
