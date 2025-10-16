package styles

func Static() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(positionProp): "static",
		}
	}
}

func Fixed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(positionProp): "fixed",
		}
	}
}

func Absolute() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(positionProp): "absolute",
		}
	}
}

func Relative() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(positionProp): "relative",
		}
	}
}

func Sticky() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(positionProp): "sticky",
		}
	}
}
