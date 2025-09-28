package styles

func HyphensNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStyleProp): "none",
		}
	}
}

func HyphensManual() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStyleProp): "manual",
		}
	}
}

func HyphensAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStyleProp): "auto",
		}
	}
}
