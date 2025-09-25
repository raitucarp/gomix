package styles

func BoxDecorationClone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxDecorationProp): "clone",
		}
	}
}

func BoxDecorationSlice() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxDecorationProp): "slice",
		}
	}
}
