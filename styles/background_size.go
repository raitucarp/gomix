package styles

import "fmt"

func BgAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "auto",
		}
	}
}

func BgCover() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "cover",
		}
	}
}

func BgContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "contain",
		}
	}
}

func BgSizeCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func BgSizeValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): fmt.Sprintf("var(--%s)", value.String()),
		}
	}
}
