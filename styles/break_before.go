package styles

func BreakBeforeAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "auto",
		}
	}
}

func BreakBeforeAvoid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "avoid",
		}
	}
}

func BreakBeforeAll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "all",
		}
	}
}

func BreakBeforeAvoidPage() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "avoid-page",
		}
	}
}

func BreakBeforePage() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "page",
		}
	}
}

func BreakBeforeLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "left",
		}
	}
}

func BreakBeforeRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "right",
		}
	}
}

func BreakBeforeColumn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakBeforeProp): "column",
		}
	}
}
