package styles

func TextLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "left",
		}
	}
}

func TextCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "center",
		}
	}
}

func TextRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "right",
		}
	}
}

func TextJustify() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "justify",
		}
	}
}

func TextStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "start",
		}
	}
}

func TextEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textAlignProp): "end",
		}
	}
}
