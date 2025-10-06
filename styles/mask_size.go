package styles

import "github.com/raitucarp/gomix/value"

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

func MaskSize(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskSizeProp): val.Value(),
		}
	}
}
