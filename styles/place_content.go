package styles

func PlaceContentCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "center",
		}
	}
}

func PlaceContentCenterSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "safe center",
		}
	}
}
func PlaceContentStart() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "start",
		}
	}
}

func PlaceContentEnd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "end",
		}
	}
}

func PlaceContentEndSafe() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "safe end",
		}
	}
}

func PlaceContentBetween() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "space-between",
		}
	}
}

func PlaceContentAround() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "space-aroud",
		}
	}
}

func PlaceContentEvenly() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "space-evenly",
		}
	}
}

func PlaceContentBaseline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "baseline",
		}
	}
}

func PlaceContentStretch() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(placeContentProp): "stretch",
		}
	}
}
