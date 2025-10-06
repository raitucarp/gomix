package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func outlineOffsetValue(val ...value.Value) string {
	propValue := ""
	if len(val) == 0 {
		propValue = "1px"
	} else {

		switch v := val[0].(type) {
		case *value.Unit[int]:
			if v.LiteralValue() <= 0 {
				propValue = fmt.Sprintf("%spx", v.Value())
			} else {
				propValue = fmt.Sprintf("calc(%spx * -1)", v.Value())
			}
		case *value.LiteralValue:
			propValue = v.Value()
		case *value.CustomProperty:
			propValue = v.Value()
		}
	}

	return propValue
}

func OutlineOffset(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(outlineOffsetProp): outlineOffsetValue(val...),
		}
	}
}
