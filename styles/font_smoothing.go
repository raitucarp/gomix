package styles

func Antialiazed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(webkitFontSmoothingProp): "antialized",
			string(mozOsxFontSmoothingProp): "gray",
		}
	}
}

func SubPixelAntialiazed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(webkitFontSmoothingProp): "auto",
			string(mozOsxFontSmoothingProp): "auto",
		}
	}
}
