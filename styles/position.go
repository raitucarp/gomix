package styles

func Static() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp): "static",
		}
	}
}

func Fixed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp): "fixed",
		}
	}
}

func Absolute() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp): "absolute",
		}
	}
}

func Relative() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp): "relative",
		}
	}
}

func Sticky() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp): "sticky",
		}
	}
}
