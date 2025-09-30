package styles

func MaskRepeat() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "repeat",
		}
	}
}

func MaskRepeatX() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "repeat-x",
		}
	}
}

func MaskRepeatY() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "repeat-y",
		}
	}
}

func MaskRepeatSpace() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "space",
		}
	}
}

func MaskRepeatRound() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "round",
		}
	}
}

func MaskNoRepeat() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRepeatProp): "no-repeat",
		}
	}
}
