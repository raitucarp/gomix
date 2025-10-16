package styles

func GridFlowRow() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoFlowProp): "row",
		}
	}
}

func GridFlowCol() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoFlowProp): "column",
		}
	}
}

func GridFlowDense() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoFlowProp): "dense",
		}
	}
}

func GridFlowRowDense() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoFlowProp): "row dense",
		}
	}
}

func GridFlowColDense() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoFlowProp): "column dense",
		}
	}
}
