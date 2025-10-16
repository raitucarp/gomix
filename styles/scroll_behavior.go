package styles

func ScrollAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollMarginProp): "auto",
		}
	}
}

func ScrollSmooth() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(scrollBehaviorProp): "smooth",
		}
	}
}
