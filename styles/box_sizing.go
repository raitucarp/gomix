package styles

func BoxBorder() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(boxSizingProp): "border-box",
		}
	}
}

func BoxContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(boxSizingProp): "content-box",
		}
	}
}
