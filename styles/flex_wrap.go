package styles

func FlexNoWrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexWrapProp): "nowrap",
		}
	}
}

func FlexWrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexWrapProp): "wrap",
		}
	}
}

func FlexWrapReverse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexWrapProp): "wrap-reverse",
		}
	}
}
