package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func GridRows(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateRowsProp): fmt.Sprintf("repeat(%d, minmax(0, 1fr))", number),
		}
	}
}

func GridRowsNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateRowsProp): "none",
		}
	}
}

func GridRowsSubgrid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateRowsProp): "subgrid",
		}
	}
}

func GridRowsBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridTemplateRowsProp): val.Value(),
		}
	}
}
