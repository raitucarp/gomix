package styles

import (
	"strconv"
)

func Z(index int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(zIndexProp): strconv.Itoa(index),
		}
	}
}

func ZAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(zIndexProp): "auto",
		}
	}
}

func ZBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(zIndexProp): value.Value(),
		}
	}
}
