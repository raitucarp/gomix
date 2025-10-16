package styles

func TextLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "left",
		}
	}
}

func TextCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "center",
		}
	}
}

func TextRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "right",
		}
	}
}

func TextJustify() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "justify",
		}
	}
}

func TextStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "start",
		}
	}
}

func TextEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textAlignProp): "end",
		}
	}
}
