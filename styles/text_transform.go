package styles

func Uppercase() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textTransformProp): "uppercase",
		}
	}
}

func Lowercase() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textTransformProp): "lowercase",
		}
	}
}

func Capitalize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textTransformProp): "capitalize",
		}
	}
}

func NormalCase() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textTransformProp): "none",
		}
	}
}
