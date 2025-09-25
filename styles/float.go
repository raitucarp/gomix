package styles

func FloatRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(floatProp): "right",
		}
	}
}

func FloatLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(floatProp): "left",
		}
	}
}

func FloatStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(floatProp): "inline-start",
		}
	}
}

func FloatEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(floatProp): "inline-end",
		}
	}
}

func FloatNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(floatProp): "none",
		}
	}
}
