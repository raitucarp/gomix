package styles

func CaptionTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(captionSideProp): "top",
		}
	}
}

func CaptionBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(captionSideProp): "bottom",
		}
	}
}
