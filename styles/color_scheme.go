package styles

func SchemeNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "normal",
		}
	}
}

func SchemeDark() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "dark",
		}
	}
}

func SchemeLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "light",
		}
	}
}

func SchemeLightDark() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "light dark",
		}
	}
}

func SchemeOnlyDark() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "only dark",
		}
	}
}

func SchemeOnlyLight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(colorSchemeProp): "only light",
		}
	}
}
