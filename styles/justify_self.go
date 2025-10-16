package styles

func JustifySelfAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "auto",
		}
	}
}

func JustifySelfStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "start",
		}
	}
}

func JustifySelfCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "center",
		}
	}
}

func JustifySelfCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "safe center",
		}
	}
}

func JustifySelfEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "end",
		}
	}
}

func JustifySelfEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "safe end",
		}
	}
}

func JustifySelfStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifySelfProp): "stretch",
		}
	}
}
