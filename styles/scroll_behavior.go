package styles

func ScrollAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollMarginProp): "auto",
		}
	}
}

func ScrollSmooth() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollBehaviorProp): "smooth",
		}
	}
}
