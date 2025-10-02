package styles

func AppearanceNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(appearanceProp): "none",
		}
	}
}

func AppearanceAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(appearanceProp): "auto",
		}
	}
}
