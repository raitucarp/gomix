package styles

func MaskAlpha() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskModeProp): "alpha",
		}
	}
}

func MaskLuminance() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskModeProp): "luminance",
		}
	}
}

func MaskMatch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskModeProp): "match-source",
		}
	}
}
