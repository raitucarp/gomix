package styles

func ResizeNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(resizeProp): "none",
		}
	}
}

func Resize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(resizeProp): "both",
		}
	}
}

func ResizeY() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(resizeProp): "vertical",
		}
	}
}

func ResizeX() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(resizeProp): "horizontal",
		}
	}
}
