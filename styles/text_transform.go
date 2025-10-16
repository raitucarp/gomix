package styles

func Uppercase() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textTransformProp): "uppercase",
		}
	}
}

func Lowercase() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textTransformProp): "lowercase",
		}
	}
}

func Capitalize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textTransformProp): "capitalize",
		}
	}
}

func NormalCase() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textTransformProp): "none",
		}
	}
}
