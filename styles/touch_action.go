package styles

func TouchAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "auto",
		}
	}
}

func TouchNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "none",
		}
	}
}

func TouchPanX() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-x",
		}
	}
}

func TouchPanLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-left",
		}
	}
}

func TouchPanRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-right",
		}
	}
}

func TouchPanY() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-y",
		}
	}
}

func TouchPanUp() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-up",
		}
	}
}

func TouchPanDown() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pan-down",
		}
	}
}

func TouchPinchZoom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "pinch-zoom",
		}
	}
}

func TouchManipulation() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(touchActionProp): "manipulation",
		}
	}
}
