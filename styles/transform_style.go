package styles

func Transform3d() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformStyleProp): "preserve-3d",
		}
	}
}

func TransformFkat() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(transformStyleProp): "flat",
		}
	}
}
