package styles

func BgRepeat() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "repeat",
		}
	}
}

func BgRepeatX() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "repeat-x",
		}
	}
}

func BgRepeatY() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "repeat-y",
		}
	}
}

func BgRepeatSpace() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "space",
		}
	}
}

func BgRepeatRound() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "round",
		}
	}
}

func BgNoRepeat() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundRepeatProp): "no-repeat",
		}
	}
}
