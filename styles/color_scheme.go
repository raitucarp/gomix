package styles

func SchemeNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "normal",
		}
	}
}

func SchemeDark() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "dark",
		}
	}
}

func SchemeLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "light",
		}
	}
}

func SchemeLightDark() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "light dark",
		}
	}
}

func SchemeOnlyDark() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "only dark",
		}
	}
}

func SchemeOnlyLight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorSchemeProp): "only light",
		}
	}
}
