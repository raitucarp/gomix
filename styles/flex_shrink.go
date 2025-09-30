package styles

func Shrink() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): "1",
		}
	}
}

func ShrinkBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexShrinkProp): value.Value(),
		}
	}
}
