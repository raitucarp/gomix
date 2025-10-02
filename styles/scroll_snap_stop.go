package styles

func SnapNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapStopProp): "normal",
		}
	}
}

func SnapAlways() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(scrollSnapStopProp): "always",
		}
	}
}
