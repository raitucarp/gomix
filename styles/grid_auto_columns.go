package styles

func AutoColsAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoColumnsProp): "auto",
		}
	}
}

func AutoColsMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoColumnsProp): "min-content",
		}
	}
}

func AutoColsMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoColumnsProp): "max-content",
		}
	}
}

func AutoColsFr() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoColumnsProp): "min-max(0,1fr)",
		}
	}
}

func AutoColsBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoColumnsProp): value.Value(),
		}
	}
}
