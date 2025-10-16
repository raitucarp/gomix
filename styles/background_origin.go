package styles

func BgOriginBorder() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundOriginProp): "border-box",
		}
	}
}

func BgOriginPadding() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundOriginProp): "padding-box",
		}
	}
}

func BgOriginContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundOriginProp): "content-box",
		}
	}
}
