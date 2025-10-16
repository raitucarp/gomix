package styles

func PlaceContentCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "center",
		}
	}
}

func PlaceContentCenterSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "safe center",
		}
	}
}
func PlaceContentStart() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "start",
		}
	}
}

func PlaceContentEnd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "end",
		}
	}
}

func PlaceContentEndSafe() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "safe end",
		}
	}
}

func PlaceContentBetween() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "space-between",
		}
	}
}

func PlaceContentAround() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "space-aroud",
		}
	}
}

func PlaceContentEvenly() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "space-evenly",
		}
	}
}

func PlaceContentBaseline() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "baseline",
		}
	}
}

func PlaceContentStretch() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(placeContentProp): "stretch",
		}
	}
}
