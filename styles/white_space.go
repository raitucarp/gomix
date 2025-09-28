package styles

func WhitespaceNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "normal",
		}
	}
}

func WhitespaceNowrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "nowrap",
		}
	}
}

func WhitespacePre() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "pre",
		}
	}
}

func WhitespacePreLine() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "pre-line",
		}
	}
}

func WhitespacePreWrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "pre-wrap",
		}
	}
}

func WhitespaceBreakSpaces() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(whiteSpaceProp): "break-spaces",
		}
	}
}
