package styles

import "github.com/raitucarp/gomix/value"

func BgAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundSizeProp): "auto",
		}
	}
}

func BgCover() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundSizeProp): "cover",
		}
	}
}

func BgContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundSizeProp): "contain",
		}
	}
}

func BgSize(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundSizeProp): val.Value(),
		}
	}
}
