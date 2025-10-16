package styles

func SnapNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapStopProp): "normal",
		}
	}
}

func SnapAlways() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollSnapStopProp): "always",
		}
	}
}
