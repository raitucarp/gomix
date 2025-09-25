package styles

func JustifySelfAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "auto",
		}
	}
}

func JustifySelfStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "start",
		}
	}
}

func JustifySelfCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "center",
		}
	}
}

func JustifySelfCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "safe center",
		}
	}
}

func JustifySelfEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "end",
		}
	}
}

func JustifySelfEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "safe end",
		}
	}
}

func JustifySelfStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifySelfProp): "stretch",
		}
	}
}
