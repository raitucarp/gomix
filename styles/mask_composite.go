package styles

func MaskAdd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskCompositeProp): "add",
		}
	}
}

func MaskSubtract() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskCompositeProp): "subtract",
		}
	}
}

func MaskIntersect() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskExclude() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskCompositeProp): "exclude",
		}
	}
}
