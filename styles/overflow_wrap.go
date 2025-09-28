package styles

func WrapBreakWord() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowWrapProp): "break-word",
		}
	}
}

func WrapAnywhere() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowWrapProp): "anywhere",
		}
	}
}

func WrapNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowWrapProp): "normal",
		}
	}
}
