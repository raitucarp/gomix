package styles

func MaskOriginBorder() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "border-box",
		}
	}
}

func MaskOriginPadding() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "padding-box",
		}
	}
}

func MaskOriginContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "content-box",
		}
	}
}

func MaskOriginFill() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "fill-box",
		}
	}
}

func MaskOriginStroke() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "stroke-box",
		}
	}
}

func MaskOriginView() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskOriginProp): "view-box",
		}
	}
}
