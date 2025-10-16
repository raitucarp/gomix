package styles

func SelfAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "auto",
		}
	}
}

func SelfStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "flex-start",
		}
	}
}

func SelfEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "flex-end",
		}
	}
}

func SelfEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "safe flex-end",
		}
	}
}

func SelfCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "center",
		}
	}
}

func SelfCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "safe center",
		}
	}
}

func SelfStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "stretch",
		}
	}
}

func SelfBaseline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "baseline",
		}
	}
}

func SelfBaselineLast() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignSelfProp): "last baseline",
		}
	}
}
