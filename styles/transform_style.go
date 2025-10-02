package styles

func Transform3d() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformStyleProp): "preserve-3d",
		}
	}
}

func TransformFkat() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transformStyleProp): "flat",
		}
	}
}
