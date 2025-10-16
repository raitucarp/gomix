package styles

func MaskTypeAlpha() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskTypeProp): "alpha",
		}
	}
}

func MaskTypeLuminance() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskTypeProp): "luminance",
		}
	}
}
