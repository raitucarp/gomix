package styles

func ObjectContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectFitProp): "contain",
		}
	}
}

func ObjectCover() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectFitProp): "cover",
		}
	}
}

func ObjectFill() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectFitProp): "fill",
		}
	}
}

func ObjectNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectFitProp): "none",
		}
	}
}

func ObjectScaleDown() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(objectFitProp): "scale-down",
		}
	}
}
