package styles

func ListInside() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStylePositionProp): "inside",
		}
	}
}

func ListOutside() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStylePositionProp): "outside",
		}
	}
}
