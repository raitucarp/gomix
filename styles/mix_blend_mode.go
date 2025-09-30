package styles

func MixBlendNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "normal",
		}
	}
}

func MixBlendMultiply() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "multiply",
		}
	}
}

func MixBlendScreen() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "screen",
		}
	}
}

func MixBlendOverlay() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "overlay",
		}
	}
}

func MixBlendDarken() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "darken",
		}
	}
}

func MixBlendLighten() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "lighten",
		}
	}
}

func MixBlendColorDodge() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "color-dodge",
		}
	}
}

func MixBlendColorBurn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "color-burn",
		}
	}
}

func MixBlendHardLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "hard-light",
		}
	}
}

func MixBlendSoftLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "soft-light",
		}
	}
}

func MixBlendDifference() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "difference",
		}
	}
}

func MixBlendExclusion() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "exclusion",
		}
	}
}

func MixBlendHue() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "hue",
		}
	}
}

func MixBlendSaturation() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "saturation",
		}
	}
}

func MixBlendColor() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "color",
		}
	}
}

func MixBlendLuminosity() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "luminosity",
		}
	}
}

func MixBlendPlusDarker() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "plus-darker",
		}
	}
}

func MixBlendPlusLighter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(mixBlendModeProp): "plus-lighter",
		}
	}
}
