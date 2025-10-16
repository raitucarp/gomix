package styles

func FloatRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(floatProp): "right",
		}
	}
}

func FloatLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(floatProp): "left",
		}
	}
}

func FloatStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(floatProp): "inline-start",
		}
	}
}

func FloatEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(floatProp): "inline-end",
		}
	}
}

func FloatNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(floatProp): "none",
		}
	}
}
