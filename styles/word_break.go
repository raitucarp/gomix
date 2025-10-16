package styles

func BreakNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(wordBreakProp): "normal",
		}
	}
}

func BreakAll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(wordBreakProp): "break-all",
		}
	}
}

func BreakKeep() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(wordBreakProp): "keep-all",
		}
	}
}
