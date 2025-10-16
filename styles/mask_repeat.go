package styles

func MaskRepeat() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "repeat",
		}
	}
}

func MaskRepeatX() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "repeat-x",
		}
	}
}

func MaskRepeatY() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "repeat-y",
		}
	}
}

func MaskRepeatSpace() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "space",
		}
	}
}

func MaskRepeatRound() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "round",
		}
	}
}

func MaskNoRepeat() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRepeatProp): "no-repeat",
		}
	}
}
