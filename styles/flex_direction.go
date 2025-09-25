package styles

func FlexRow() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexDirectionProp): "row",
		}
	}
}

func FlexRowReverse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexDirectionProp): "row-reverse",
		}
	}
}

func FlexCol() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexDirectionProp): "column",
		}
	}
}

func FlexColReverse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(flexDirectionProp): "column-reverse",
		}
	}
}
