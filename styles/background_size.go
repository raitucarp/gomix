package styles

func BgAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "auto",
		}
	}
}

func BgCover() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "cover",
		}
	}
}

func BgContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): "contain",
		}
	}
}

func BgSize(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backgroundSizeProp): value.Value(),
		}
	}
}
