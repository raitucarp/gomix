package styles

func WrapBreakWord() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowWrapProp): "break-word",
		}
	}
}

func WrapAnywhere() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowWrapProp): "anywhere",
		}
	}
}

func WrapNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overflowWrapProp): "normal",
		}
	}
}
