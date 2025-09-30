package styles

func MaskClipBorder() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "border-box",
		}
	}
}

func MaskClipPadding() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "padding-box",
		}
	}
}

func MaskClipContent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "content-box",
		}
	}
}

func MaskClipFill() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "fill-box",
		}
	}
}

func MaskClipStroke() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "stroke-box",
		}
	}
}

func MaskClipView() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "view-box",
		}
	}
}

func MaskNoClip() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskClipProp): "no-clip",
		}
	}
}
