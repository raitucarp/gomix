package styles

func BackfaceHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backfaceVisibilityProp): "hidden",
		}
	}
}

func BackfaceVisible() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backfaceVisibilityProp): "visible",
		}
	}
}
