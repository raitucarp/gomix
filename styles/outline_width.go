package styles

import "github.com/raitucarp/gomix/value"

func Outline(val ...value.Value) ApplyProp {
	return func(s *Style) StyleProp {

		return &Properties{
			string(outlineWidthProp): propWidthValue(val...),
		}
	}
}
