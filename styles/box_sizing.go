package styles

func BoxBorder() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxSizingProp): "border-box",
		}
	}
}

func BoxContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxSizingProp): "content-box",
		}
	}
}
