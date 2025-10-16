package styles

func ItemsStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "flex-start",
		}
	}
}

func ItemsEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "flex-end",
		}
	}
}

func ItemsEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "safe flex-end",
		}
	}
}

func ItemsCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "center",
		}
	}
}

func ItemsCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "safe center",
		}
	}
}

func ItemsBaseline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "baseline",
		}
	}
}

func ItemsBaselineLast() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "last baseline",
		}
	}
}

func ItemsStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignItemsProp): "stretch",
		}
	}
}
