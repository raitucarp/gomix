package styles

func BgFixed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundAttachmentProp): "fixed",
		}
	}
}

func BgLocal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundAttachmentProp): "fixed",
		}
	}
}

func BgScroll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundAttachmentProp): "scroll",
		}
	}
}
