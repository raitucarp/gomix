package styles

func SnapStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapAlignProp): "start",
		}
	}
}

func SnapEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapAlignProp): "end",
		}
	}
}

func SnapCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapAlignProp): "center",
		}
	}
}

func SnapAlignNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapAlignProp): "none",
		}
	}
}
