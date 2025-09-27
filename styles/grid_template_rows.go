package styles

import (
	"fmt"
)

func GridRows(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateRowsProp): fmt.Sprintf("repeat(%d, minmax(0, 1fr))", number),
		}
	}
}

func GridRowsNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateRowsProp): "none",
		}
	}
}

func GridRowsSubgrid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateRowsProp): "subgrid",
		}
	}
}

func GridRowsCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateRowsProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func GridRowsValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateRowsProp): value.String(),
		}
	}
}
