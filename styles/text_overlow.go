package styles

func Truncate() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):     "hidden",
			string(textOverflowProp): "ellipsis",
			string(whiteSpaceProp):   "nowrap",
		}
	}
}

func TextEllipsis() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textOverflowProp): "ellipsis",
		}
	}
}

func TextClip() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textOverflowProp): "clip",
		}
	}
}
