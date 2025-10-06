package styles

import (
	"strconv"

	"github.com/raitucarp/gomix/value"
)

func LineClamp(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):        "hidden",
			string(displayProp):         "-webkit-box",
			string(webkitBoxOrientProp): "vertical",
			string(webkitLineClampProp): strconv.Itoa(number),
		}
	}
}

func LineClampNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):        "visible",
			string(displayProp):         "block",
			string(webkitBoxOrientProp): "horizontal",
			string(webkitLineClampProp): "unset",
		}
	}
}

func LineClampBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):        "hidden",
			string(displayProp):         "-webkit-box",
			string(webkitBoxOrientProp): "vertical",
			string(webkitLineClampProp): val.Value(),
		}
	}
}
