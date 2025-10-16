package styles

func BoxDecorationClone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(boxDecorationProp): "clone",
		}
	}
}

func BoxDecorationSlice() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(boxDecorationProp): "slice",
		}
	}
}
