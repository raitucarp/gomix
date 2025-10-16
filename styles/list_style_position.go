package styles

func ListInside() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStylePositionProp): "inside",
		}
	}
}

func ListOutside() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(listStylePositionProp): "outside",
		}
	}
}
