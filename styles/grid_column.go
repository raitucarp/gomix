package styles

import "strconv"

func ColSpan(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridColumnProp): strconv.Itoa(number),
		}
	}
}
