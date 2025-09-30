package styles

func ListDisc() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "disc",
		}
	}
}

func ListDecimal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "decimal",
		}
	}
}

func ListNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): "none",
		}
	}
}

func List(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleTypeProp): value.Value(),
		}
	}
}
