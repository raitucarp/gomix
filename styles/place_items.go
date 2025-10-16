package styles

func PlaceItemsStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "start",
		}
	}
}

func PlaceItemsEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "end",
		}
	}
}

func PlaceItemsEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "safe end",
		}
	}
}

func PlaceItemsCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "center",
		}
	}
}

func PlaceItemsCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "safe center",
		}
	}
}

func PlaceItemsBaseline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "baseline",
		}
	}
}

func PlaceItemsStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeItemsProp): "stretch",
		}
	}
}
