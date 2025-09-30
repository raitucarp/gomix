package styles

func Outline(value ...customValue) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(outlineWidthProp): propWidthValue(value...),
		}
	}
}
