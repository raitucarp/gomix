package styles

func DecorationSolid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorStyleProp): "solid",
		}
	}
}

func DecorationDouble() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorStyleProp): "double",
		}
	}
}

func DecorationDotted() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorStyleProp): "dotted",
		}
	}
}

func DecorationDashed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorStyleProp): "dashed",
		}
	}
}

func DecorationWavy() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textDecorStyleProp): "wavy",
		}
	}
}
