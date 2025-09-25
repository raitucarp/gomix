package styles

func PlaceSelfAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "auto",
		}
	}
}

func PlaceSelfStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "start",
		}
	}
}

func PlaceSelfEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "end",
		}
	}
}

func PlaceSelfEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "safe end",
		}
	}
}

func PlaceSelfCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "center",
		}
	}
}

func PlaceSelfCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "safe center",
		}
	}
}

func PlaceSelfStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeSelfProp): "stretch",
		}
	}
}
