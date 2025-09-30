package styles

func MaskTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "top left",
		}
	}
}

func MaskTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "top",
		}
	}

}

func MaskTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "top right",
		}
	}
}

func MaskLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "left",
		}
	}
}

func MaskCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "center",
		}
	}
}

func MaskRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "right",
		}
	}
}

func MaskBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "bottom left",
		}
	}
}

func MaskBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "bottom",
		}
	}
}

func MaskBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): "bottom right",
		}
	}
}

func MaskPosition(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskPositionProp): value.Value(),
		}
	}
}
