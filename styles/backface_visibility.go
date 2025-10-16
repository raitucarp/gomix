package styles

func BackfaceHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backfaceVisibilityProp): "hidden",
		}
	}
}

func BackfaceVisible() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backfaceVisibilityProp): "visible",
		}
	}
}
