package styles

func GridFlowRow() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoFlowProp): "row",
		}
	}
}

func GridFlowCol() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoFlowProp): "column",
		}
	}
}

func GridFlowDense() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoFlowProp): "dense",
		}
	}
}

func GridFlowRowDense() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoFlowProp): "row dense",
		}
	}
}

func GridFlowColDense() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoFlowProp): "column dense",
		}
	}
}
