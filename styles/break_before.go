package styles

func BreakBeforeAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "auto",
		}
	}
}

func BreakBeforeAvoid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "avoid",
		}
	}
}

func BreakBeforeAll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "all",
		}
	}
}

func BreakBeforeAvoidPage() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "avoid-page",
		}
	}
}

func BreakBeforePage() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "page",
		}
	}
}

func BreakBeforeLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "left",
		}
	}
}

func BreakBeforeRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "right",
		}
	}
}

func BreakBeforeColumn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakBeforeProp): "column",
		}
	}
}
