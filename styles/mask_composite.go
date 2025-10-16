package styles

func MaskAdd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskCompositeProp): "add",
		}
	}
}

func MaskSubtract() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskCompositeProp): "subtract",
		}
	}
}

func MaskIntersect() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskExclude() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskCompositeProp): "exclude",
		}
	}
}
