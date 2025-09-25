package styles

func BreakAfterAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "auto",
		}
	}
}

func BreakAfterAvoid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "avoid",
		}
	}
}

func BreakAfterAll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "all",
		}
	}
}

func BreakAfterAvoidPage() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "avoid-page",
		}
	}
}

func BreakAfterPage() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "page",
		}
	}
}

func BreakAfterLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "left",
		}
	}
}

func BreakAfterRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "right",
		}
	}
}

func BreakAfterColumn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakAfterProp): "column",
		}
	}
}
