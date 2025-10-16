package styles

func SelectNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(userSelectProp): "none",
		}
	}
}

func SelectText() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(userSelectProp): "text",
		}
	}
}

func SelectAll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(userSelectProp): "all",
		}
	}
}

func SelectAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(userSelectProp): "auto",
		}
	}
}
