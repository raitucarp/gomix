package styles

func BgRepeat() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "repeat",
		}
	}
}

func BgRepeatX() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "repeat-x",
		}
	}
}

func BgRepeatY() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "repeat-y",
		}
	}
}

func BgRepeatSpace() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "space",
		}
	}
}

func BgRepeatRound() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "round",
		}
	}
}

func BgNoRepeat() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(backgroundRepeatProp): "no-repeat",
		}
	}
}
