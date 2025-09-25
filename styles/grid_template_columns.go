package styles

import (
	"fmt"
	"strconv"
)

func GridCols(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): fmt.Sprintf("repeat(%s, minmax(0, 1fr))", strconv.Itoa(number)),
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
