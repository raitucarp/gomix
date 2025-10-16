package styles

func MixBlendNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "normal",
		}
	}
}

func MixBlendMultiply() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "multiply",
		}
	}
}

func MixBlendScreen() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "screen",
		}
	}
}

func MixBlendOverlay() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "overlay",
		}
	}
}

func MixBlendDarken() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "darken",
		}
	}
}

func MixBlendLighten() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "lighten",
		}
	}
}

func MixBlendColorDodge() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "color-dodge",
		}
	}
}

func MixBlendColorBurn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "color-burn",
		}
	}
}

func MixBlendHardLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "hard-light",
		}
	}
}

func MixBlendSoftLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "soft-light",
		}
	}
}

func MixBlendDifference() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "difference",
		}
	}
}

func MixBlendExclusion() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "exclusion",
		}
	}
}

func MixBlendHue() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "hue",
		}
	}
}

func MixBlendSaturation() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "saturation",
		}
	}
}

func MixBlendColor() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "color",
		}
	}
}

func MixBlendLuminosity() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "luminosity",
		}
	}
}

func MixBlendPlusDarker() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "plus-darker",
		}
	}
}

func MixBlendPlusLighter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(mixBlendModeProp): "plus-lighter",
		}
	}
}
