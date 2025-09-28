package styles

import (
	"fmt"
	"strconv"
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

func LineClampCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):        "hidden",
			string(displayProp):         "-webkit-box",
			string(webkitBoxOrientProp): "vertical",
			string(webkitLineClampProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func LineClampValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(overflowProp):        "hidden",
			string(displayProp):         "-webkit-box",
			string(webkitBoxOrientProp): "vertical",
			string(webkitLineClampProp): value.String(),
		}
	}
}
