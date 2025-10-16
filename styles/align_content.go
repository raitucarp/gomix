package styles

func ContentNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "normal",
		}
	}
}

func ContentCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "center",
		}
	}
}

func ContentStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "flex-start",
		}
	}
}

func ContentEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "flex-end",
		}
	}
}

func ContentBetween() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "space-between",
		}
	}
}

func ContentAround() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "space-around",
		}
	}
}

func ContentEvenly() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "space-evenly",
		}
	}
}

func ContentBaseline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "baseline",
		}
	}
}

func ContentStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(alignContentProp): "stretch",
		}
	}
}
