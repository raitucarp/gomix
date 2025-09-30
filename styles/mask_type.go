package styles

func MaskTypeAlpha() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskTypeProp): "alpha",
		}
	}
}

func MaskTypeLuminance() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskTypeProp): "luminance",
		}
	}
}
