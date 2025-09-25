package styles

func ObjectContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectFitProp): "contain",
		}
	}
}

func ObjectCover() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectFitProp): "cover",
		}
	}
}

func ObjectFill() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectFitProp): "fill",
		}
	}
}

func ObjectNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectFitProp): "none",
		}
	}
}

func ObjectScaleDown() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectFitProp): "scale-down",
		}
	}
}
