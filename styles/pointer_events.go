package styles

func PointerEventsAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(pointerEventsProp): "fixed",
		}
	}
}

func PointerEventsNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(pointerEventsProp): "none",
		}
	}
}
