package styles

func FlexRow() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexDirectionProp): "row",
		}
	}
}

func FlexRowReverse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexDirectionProp): "row-reverse",
		}
	}
}

func FlexCol() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexDirectionProp): "column",
		}
	}
}

func FlexColReverse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(flexDirectionProp): "column-reverse",
		}
	}
}
