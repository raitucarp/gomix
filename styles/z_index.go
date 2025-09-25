package styles

import (
	"fmt"
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

func ZCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(zIndexProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}
