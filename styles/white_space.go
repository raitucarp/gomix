package styles

func WhitespaceNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "normal",
		}
	}
}

func WhitespaceNowrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "nowrap",
		}
	}
}

func WhitespacePre() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "pre",
		}
	}
}

func WhitespacePreLine() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "pre-line",
		}
	}
}

func WhitespacePreWrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "pre-wrap",
		}
	}
}

func WhitespaceBreakSpaces() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(whiteSpaceProp): "break-spaces",
		}
	}
}
