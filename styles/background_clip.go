package styles

func BgClipBorder() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundClipProp): "border-box",
		}
	}
}

func BgClipPadding() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundClipProp): "padding-box",
		}
	}
}

func BgClipContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundClipProp): "content-box",
		}
	}
}

func BgClipText() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundClipProp): "text",
		}
	}
}
