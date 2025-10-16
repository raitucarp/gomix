package styles

func DecorationSolid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorStyleProp): "solid",
		}
	}
}

func DecorationDouble() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorStyleProp): "double",
		}
	}
}

func DecorationDotted() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorStyleProp): "dotted",
		}
	}
}

func DecorationDashed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorStyleProp): "dashed",
		}
	}
}

func DecorationWavy() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorStyleProp): "wavy",
		}
	}
}
