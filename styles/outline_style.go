package styles

func OutlineSolid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineStyleProp): "solid",
		}
	}
}

func OutlineDashed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineStyleProp): "dashed",
		}
	}
}

func OutlineDotted() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineStyleProp): "dotted",
		}
	}
}

func OutlineDouble() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineStyleProp): "double",
		}
	}
}

func OutlineHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineProp):      "2px solid transparent",
			string(outlineStyleProp): "2px",
		}
	}
}

func OutlineNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(outlineStyleProp): "none",
		}
	}
}
