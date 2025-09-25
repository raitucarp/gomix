package styles

func FlexNoWrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexWrapProp): "nowrap",
		}
	}
}

func FlexWrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexWrapProp): "wrap",
		}
	}
}

func FlexWrapReverse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexWrapProp): "wrap-reverse",
		}
	}
}
