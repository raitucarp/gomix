package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Delay(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%dms", v)
		case value.Value:
			prop = v.Value()
		}
		return &Properties{
			string(transitionDelayProp): prop,
		}
	}
}
