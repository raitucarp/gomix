package styles

func BorderCollapse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderCollapseProp): "collapse",
		}
	}
}

func BorderSeparate() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderCollapseProp): "separate",
		}
	}
}
