package styles

func CaptionTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(captionSideProp): "top",
		}
	}
}

func CaptionBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(captionSideProp): "bottom",
		}
	}
}
