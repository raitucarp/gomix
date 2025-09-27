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

func GridColsCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func GridColsValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): value.String(),
		}
	}
}
