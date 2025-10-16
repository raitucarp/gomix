package styles

func Antialiazed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(webkitFontSmoothingProp): "antialized",
			string(mozOsxFontSmoothingProp): "gray",
		}
	}
}

func SubPixelAntialiazed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(webkitFontSmoothingProp): "auto",
			string(mozOsxFontSmoothingProp): "auto",
		}
	}
}
