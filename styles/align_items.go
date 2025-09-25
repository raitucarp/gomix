package styles

func ItemsStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "flex-start",
		}
	}
}

func ItemsEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "flex-end",
		}
	}
}

func ItemsEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "safe flex-end",
		}
	}
}

func ItemsCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "center",
		}
	}
}

func ItemsCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "safe center",
		}
	}
}

func ItemsBaseline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "baseline",
		}
	}
}

func ItemsBaselineLast() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "last baseline",
		}
	}
}

func ItemsStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(alignItemsProp): "stretch",
		}
	}
}
