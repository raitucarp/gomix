package styles

import "github.com/raitucarp/gomix/value"

func MaskAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskSizeProp): "auto",
		}
	}
}

func MaskCover() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskSizeProp): "cover",
		}
	}
}

func MaskContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskSizeProp): "contain",
		}
	}
}

func MaskSize(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskSizeProp): val.Value(),
		}
	}
}
