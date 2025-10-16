package styles

func PlaceSelfAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "auto",
		}
	}
}

func PlaceSelfStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "start",
		}
	}
}

func PlaceSelfEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "end",
		}
	}
}

func PlaceSelfEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "safe end",
		}
	}
}

func PlaceSelfCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "center",
		}
	}
}

func PlaceSelfCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "safe center",
		}
	}
}

func PlaceSelfStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeSelfProp): "stretch",
		}
	}
}
