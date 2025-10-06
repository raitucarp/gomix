package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Stroke(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%v", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(strokeWidthProp): prop,
		}
	}
}
