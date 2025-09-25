package styles

func ClearLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "left",
		}
	}
}

func ClearRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "right",
		}
	}
}

func ClearBoth() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "both",
		}
	}
}

func ClearStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "inline-start",
		}
	}
}

func ClearEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "inline-end",
		}
	}
}

func ClearNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "none",
		}
	}
}
