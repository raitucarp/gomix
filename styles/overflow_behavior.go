package styles

func OverscrollAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorProp): "auto",
		}
	}
}

func OverscrollContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorProp): "contain",
		}
	}
}

func OverscrollNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorProp): "none",
		}
	}
}

func OverscrollXAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorXProp): "auto",
		}
	}
}

func OverscrollXContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorXProp): "contain",
		}
	}
}

func OverscrollXNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorXProp): "none",
		}
	}
}

func OverscrollYAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorYProp): "auto",
		}
	}
}

func OverscrollYContain() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorYProp): "contain",
		}
	}
}

func OverscrollYNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(overscrollBehaviorYProp): "none",
		}
	}
}
