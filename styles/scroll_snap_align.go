package styles

func SnapStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapAlignProp): "start",
		}
	}
}

func SnapEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapAlignProp): "end",
		}
	}
}

func SnapCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapAlignProp): "center",
		}
	}
}

func SnapAlignNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapAlignProp): "none",
		}
	}
}
