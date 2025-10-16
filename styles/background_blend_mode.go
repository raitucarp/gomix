package styles

func BgBlendNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "normal",
		}
	}
}

func BgBlendMultiply() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "multiply",
		}
	}
}

func BgBlendScreen() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "screen",
		}
	}
}

func BgBlendOverlay() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "overlay",
		}
	}
}

func BgBlendDarken() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "darken",
		}
	}
}

func BgBlendLighten() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "lighten",
		}
	}
}

func BgBlendColorDodge() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "color-dodge",
		}
	}
}

func BgBlendColorBurn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "color-burn",
		}
	}
}

func BgBlendHardLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "hard-light",
		}
	}
}

func BgBlendSoftLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "soft-light",
		}
	}
}

func BgBlendDifference() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "difference",
		}
	}
}

func BgBlendExclusion() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "exclusion",
		}
	}
}

func BgBlendHue() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "hue",
		}
	}
}

func BgBlendSaturation() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "saturation",
		}
	}
}

func BgBlendColor() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "color",
		}
	}
}

func BgBlendLuminosity() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundBlendModeProp): "luminosity",
		}
	}
}
