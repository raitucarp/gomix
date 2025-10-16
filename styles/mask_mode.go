package styles

func MaskAlpha() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskModeProp): "alpha",
		}
	}
}

func MaskLuminance() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskModeProp): "luminance",
		}
	}
}

func MaskMatch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskModeProp): "match-source",
		}
	}
}
