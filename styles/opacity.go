package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func opacityValue(val value.Value) string {
	propValue := ""
	switch v := val.(type) {
	case *value.LiteralValue:
		propValue = v.Value()
	case *value.Unit[int], *value.Unit[float32], *value.Unit[float64]:
		propValue = fmt.Sprintf("%#v%%", v)

	case *value.CustomProperty:
		propValue = v.Value()
	}

	return propValue
}

func Opacity(val value.Value) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(opacityProp): opacityValue(val),
		}
	}
}
