package styles

func MaskOriginBorder() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "border-box",
		}
	}
}

func MaskOriginPadding() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "padding-box",
		}
	}
}

func MaskOriginContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "content-box",
		}
	}
}

func MaskOriginFill() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "fill-box",
		}
	}
}

func MaskOriginStroke() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "stroke-box",
		}
	}
}

func MaskOriginView() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskOriginProp): "view-box",
		}
	}
}
