package styles

func BgBlendNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "normal",
		}
	}
}

func BgBlendMultiply() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "multiply",
		}
	}
}

func BgBlendScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "screen",
		}
	}
}

func BgBlendOverlay() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "overlay",
		}
	}
}

func BgBlendDarken() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "darken",
		}
	}
}

func BgBlendLighten() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "lighten",
		}
	}
}

func BgBlendColorDodge() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "color-dodge",
		}
	}
}

func BgBlendColorBurn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "color-burn",
		}
	}
}

func BgBlendHardLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "hard-light",
		}
	}
}

func BgBlendSoftLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "soft-light",
		}
	}
}

func BgBlendDifference() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "difference",
		}
	}
}

func BgBlendExclusion() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "exclusion",
		}
	}
}

func BgBlendHue() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "hue",
		}
	}
}

func BgBlendSaturation() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "saturation",
		}
	}
}

func BgBlendColor() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "color",
		}
	}
}

func BgBlendLuminosity() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundBlendModeProp): "luminosity",
		}
	}
}
