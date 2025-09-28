package styles

func BreakNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(wordBreakProp): "normal",
		}
	}
}

func BreakAll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(wordBreakProp): "break-all",
		}
	}
}

func BreakKeep() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(wordBreakProp): "keep-all",
		}
	}
}
