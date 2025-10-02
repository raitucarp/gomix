package styles

func FieldSizingFixed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fieldSizingPropp): "fixed",
		}
	}
}

func FieldSizingContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fieldSizingPropp): "content",
		}
	}
}
