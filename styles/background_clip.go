package styles

func BgClipBorder() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundClipProp): "border-box",
		}
	}
}

func BgClipPadding() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundClipProp): "padding-box",
		}
	}
}

func BgClipContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundClipProp): "content-box",
		}
	}
}

func BgClipText() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundClipProp): "text",
		}
	}
}
