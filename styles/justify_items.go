package styles

func JustifyItemsStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "start",
		}
	}
}

func JustifyItemsEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "end",
		}
	}
}

func JustifyItemsEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "safe end",
		}
	}
}

func JustifyItemsCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "center",
		}
	}
}

func JustifyItemsCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "safe center",
		}
	}
}

func JustifyItemsStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "stretch",
		}
	}
}

func JustifyItemsNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(justifyItemsProp): "normal",
		}
	}
}
