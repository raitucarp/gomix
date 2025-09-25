package styles

func PlaceItemsStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "start",
		}
	}
}

func PlaceItemsEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "end",
		}
	}
}

func PlaceItemsEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "safe end",
		}
	}
}

func PlaceItemsCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "center",
		}
	}
}

func PlaceItemsCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "safe center",
		}
	}
}

func PlaceItemsBaseline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "baseline",
		}
	}
}

func PlaceItemsStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeItemsProp): "stretch",
		}
	}
}
