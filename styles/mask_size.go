package styles

func MaskAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskSizeProp): "auto",
		}
	}
}

func MaskCover() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskSizeProp): "cover",
		}
	}
}

func MaskContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskSizeProp): "contain",
		}
	}
}

func MaskSize(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskSizeProp): value.Value(),
		}
	}
}
