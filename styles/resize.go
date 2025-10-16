package styles

func ResizeNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(resizeProp): "none",
		}
	}
}

func Resize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(resizeProp): "both",
		}
	}
}

func ResizeY() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(resizeProp): "vertical",
		}
	}
}

func ResizeX() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(resizeProp): "horizontal",
		}
	}
}
