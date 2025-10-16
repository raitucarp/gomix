package styles

func FieldSizingFixed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fieldSizingPropp): "fixed",
		}
	}
}

func FieldSizingContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fieldSizingPropp): "content",
		}
	}
}
