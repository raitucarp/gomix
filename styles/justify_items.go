package styles

func JustifyItemsStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "start",
		}
	}
}

func JustifyItemsEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "end",
		}
	}
}

func JustifyItemsEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "safe end",
		}
	}
}

func JustifyItemsCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "center",
		}
	}
}

func JustifyItemsCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "safe center",
		}
	}
}

func JustifyItemsStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "stretch",
		}
	}
}

func JustifyItemsNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(justifyItemsProp): "normal",
		}
	}
}
