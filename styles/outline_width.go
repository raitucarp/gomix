package styles

import "github.com/raitucarp/gomix/value"

func Outline(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(outlineWidthProp): propWidthValue(val...),
		}
	}
}
