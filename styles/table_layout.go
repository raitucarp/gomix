package styles

func TableAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(tableLayoutProp): "auto",
		}
	}
}

func TableFixed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(tableLayoutProp): "fixed",
		}
	}
}
