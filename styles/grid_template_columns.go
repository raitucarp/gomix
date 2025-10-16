package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func GridCols(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateColumnsProp): fmt.Sprintf("repeat(%d, minmax(0, 1fr))", number),
		}
	}
}

func GridColsNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateColumnsProp): "none",
		}
	}
}

func GridColsSubgrid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateColumnsProp): "subgrid",
		}
	}
}

func GridColsBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateColumnsProp): val.Value(),
		}
	}
}
