package styles

func Underline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorLineProp): "underline",
		}
	}
}

func Overline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorLineProp): "overline",
		}
	}
}

func LineThrough() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorLineProp): "line-through",
		}
	}
}

func NoUnderline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorLineProp): "none",
		}
	}
}
