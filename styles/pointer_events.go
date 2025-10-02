package styles

func PointerEventsAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(pointerEventsProp): "fixed",
		}
	}
}

func PointerEventsNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(pointerEventsProp): "none",
		}
	}
}
