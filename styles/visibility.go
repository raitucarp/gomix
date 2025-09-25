package styles

func Visible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(visibilityProp): "visible",
		}
	}
}

func Invisible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(visibilityProp): "hidden",
		}
	}
}

func Collapse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(visibilityProp): "collapse",
		}
	}
}
