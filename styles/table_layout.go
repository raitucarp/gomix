package styles

func TableAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(tableLayoutProp): "auto",
		}
	}
}

func TableFixed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(tableLayoutProp): "fixed",
		}
	}
}
