package styles

func MaskClipBorder() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "border-box",
		}
	}
}

func MaskClipPadding() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "padding-box",
		}
	}
}

func MaskClipContent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "content-box",
		}
	}
}

func MaskClipFill() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "fill-box",
		}
	}
}

func MaskClipStroke() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "stroke-box",
		}
	}
}

func MaskClipView() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "view-box",
		}
	}
}

func MaskNoClip() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskClipProp): "no-clip",
		}
	}
}
