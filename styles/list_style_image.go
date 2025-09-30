package styles

func ListImage(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): value.Value(),
		}
	}
}

func ListImageNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): "none",
		}
	}
}
