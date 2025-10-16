package styles

func BorderSolid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "solid",
		}
	}
}

func BorderDashed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "dashed",
		}
	}
}

func BorderDotted() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "dotted",
		}
	}
}

func BorderDouble() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "double",
		}
	}
}

func BorderHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "hidden",
		}
	}
}

func BorderNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStyleProp): "none",
		}
	}
}

//--

func DivideSolid() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "solid",
		}
	}
}

func DivideDashed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "dashed",
		}
	}
}

func DivideDotted() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "dotted",
		}
	}
}

func DivideDouble() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "double",
		}
	}
}

func DivideHidden() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "hidden",
		}
	}
}

func DivideNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			notLastChildProp(borderStyleProp): "none",
		}
	}
}
