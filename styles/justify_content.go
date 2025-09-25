package styles

func JustifyStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "flex-start",
		}
	}
}

func JustifyEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "flex-end",
		}
	}
}

func JustifyEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "safe flex-end",
		}
	}
}

func JustifyCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "center",
		}
	}
}

func JustifyCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "safe center",
		}
	}
}

func JustifyBetween() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "space-between",
		}
	}
}

func JustifyAround() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "space-aroud",
		}
	}
}

func JustifyEvenly() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "space-evenly",
		}
	}
}

func JustifyStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "stretch",
		}
	}
}

func JustifyBaseline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "baseline",
		}
	}
}

func JustifyNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyContentProp): "normal",
		}
	}
}
