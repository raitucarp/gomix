package styles

func BreakInsideAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakInsideProp): "auto",
		}
	}
}

func BreakInsideAvoid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakInsideProp): "avoid",
		}
	}
}

func BreakInsideAvoidPage() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakInsideProp): "avoid-page",
		}
	}
}

func BreakInsideAvoidColumn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(breakInsideProp): "avoid-column",
		}
	}
}
