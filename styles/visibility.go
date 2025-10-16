package styles

func Visible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(visibilityProp): "visible",
		}
	}
}

func Invisible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(visibilityProp): "hidden",
		}
	}
}

func Collapse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(visibilityProp): "collapse",
		}
	}
}
