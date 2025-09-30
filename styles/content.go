package styles

func Content(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): value.Value(),
		}
	}
}

func ContentNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): "none",
		}
	}
}
