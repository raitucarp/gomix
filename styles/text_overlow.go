package styles

func Truncate() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowProp):     "hidden",
			string(textOverflowProp): "ellipsis",
			string(whiteSpaceProp):   "nowrap",
		}
	}
}

func TextEllipsis() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textOverflowProp): "ellipsis",
		}
	}
}

func TextClip() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textOverflowProp): "clip",
		}
	}
}
