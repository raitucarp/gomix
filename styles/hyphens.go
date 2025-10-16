package styles

func HyphensNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontStyleProp): "none",
		}
	}
}

func HyphensManual() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontStyleProp): "manual",
		}
	}
}

func HyphensAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontStyleProp): "auto",
		}
	}
}
