package styles

func BreakAfterAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "auto",
		}
	}
}

func BreakAfterAvoid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "avoid",
		}
	}
}

func BreakAfterAll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "all",
		}
	}
}

func BreakAfterAvoidPage() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "avoid-page",
		}
	}
}

func BreakAfterPage() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "page",
		}
	}
}

func BreakAfterLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "left",
		}
	}
}

func BreakAfterRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "right",
		}
	}
}

func BreakAfterColumn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakAfterProp): "column",
		}
	}
}
