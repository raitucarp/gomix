package styles

func Underline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorLineProp): "underline",
		}
	}
}

func Overline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorLineProp): "overline",
		}
	}
}

func LineThrough() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorLineProp): "line-through",
		}
	}
}

func NoUnderline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorLineProp): "none",
		}
	}
}
