package styles

func BreakInsideAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakInsideProp): "auto",
		}
	}
}

func BreakInsideAvoid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakInsideProp): "avoid",
		}
	}
}

func BreakInsideAvoidPage() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakInsideProp): "avoid-page",
		}
	}
}

func BreakInsideAvoidColumn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(breakInsideProp): "avoid-column",
		}
	}
}
