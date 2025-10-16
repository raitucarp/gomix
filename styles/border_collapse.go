package styles

func BorderCollapse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderCollapseProp): "collapse",
		}
	}
}

func BorderSeparate() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderCollapseProp): "separate",
		}
	}
}
