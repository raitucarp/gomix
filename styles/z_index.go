package styles

import (
	"strconv"

	"github.com/raitucarp/gomix/value"
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

func ZBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(zIndexProp): val.Value(),
		}
	}
}
