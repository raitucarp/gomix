package styles

func OverscrollAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorProp): "auto",
		}
	}
}

func OverscrollContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorProp): "contain",
		}
	}
}

func OverscrollNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorProp): "none",
		}
	}
}

func OverscrollXAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorXProp): "auto",
		}
	}
}

func OverscrollXContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorXProp): "contain",
		}
	}
}

func OverscrollXNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorXProp): "none",
		}
	}
}

func OverscrollYAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorYProp): "auto",
		}
	}
}

func OverscrollYContain() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorYProp): "contain",
		}
	}
}

func OverscrollYNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overscrollBehaviorYProp): "none",
		}
	}
}
