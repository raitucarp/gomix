package styles

func BgOriginBorder() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundOriginProp): "border-box",
		}
	}
}

func BgOriginPadding() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundOriginProp): "padding-box",
		}
	}
}

func BgOriginContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundOriginProp): "content-box",
		}
	}
}
