package styles

import (
	"fmt"
)

func GridCols(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): fmt.Sprintf("repeat(%d, minmax(0, 1fr))", number),
		}
	}
}

func GridColsNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): "none",
		}
	}
}

func GridColsSubgrid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): "subgrid",
		}
	}
}

func GridColsBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): value.Value(),
		}
	}
}
