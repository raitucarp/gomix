package styles

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

func BgPosition(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundPositionProp): value.Value(),
		}
	}
}
