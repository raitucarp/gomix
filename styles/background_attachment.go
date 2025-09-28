package styles

func BgFixed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundAttachmentProp): "fixed",
		}
	}
}

func BgLocal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundAttachmentProp): "fixed",
		}
	}
}

func BgScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundAttachmentProp): "scroll",
		}
	}
}
