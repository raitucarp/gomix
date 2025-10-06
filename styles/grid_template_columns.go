package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
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

func GridColsBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridTemplateColumnsProp): val.Value(),
		}
	}
}
