package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Duration(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("%dms", v)
		case value.Value:
			prop = v.Value()
		}
		return &properties{
			string(transitionDurationProp): prop,
		}
	}
}

func DurationInitial() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionDurationProp): "initial",
		}
	}
}
