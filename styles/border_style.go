package styles

func BorderSolid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "solid",
		}
	}
}

func BorderDashed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "dashed",
		}
	}
}

func BorderDotted() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "dotted",
		}
	}
}

func BorderDouble() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "double",
		}
	}
}

func BorderHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "hidden",
		}
	}
}

func BorderNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStyleProp): "none",
		}
	}
}

//--

func DivideSolid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "solid",
		}
	}
}

func DivideDashed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "dashed",
		}
	}
}

func DivideDotted() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "dotted",
		}
	}
}

func DivideDouble() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "double",
		}
	}
}

func DivideHidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "hidden",
		}
	}
}

func DivideNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderStyleProp): "none",
		}
	}
}
