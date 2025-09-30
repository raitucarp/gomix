package styles

func OutlineSolid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineStyleProp): "solid",
		}
	}
}

func OutlineDashed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineStyleProp): "dashed",
		}
	}
}

func OutlineDotted() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineStyleProp): "dotted",
		}
	}
}

func OutlineDouble() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineStyleProp): "double",
		}
	}
}

func OutlineHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineProp):      "2px solid transparent",
			string(outlineStyleProp): "2px",
		}
	}
}

func OutlineNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineStyleProp): "none",
		}
	}
}
