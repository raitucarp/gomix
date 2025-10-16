package styles

func AppearanceNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(appearanceProp): "none",
		}
	}
}

func AppearanceAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(appearanceProp): "auto",
		}
	}
}
