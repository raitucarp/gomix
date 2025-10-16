package styles

func ClearLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "left",
		}
	}
}

func ClearRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "right",
		}
	}
}

func ClearBoth() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "both",
		}
	}
}

func ClearStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "inline-start",
		}
	}
}

func ClearEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "inline-end",
		}
	}
}

func ClearNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "none",
		}
	}
}
