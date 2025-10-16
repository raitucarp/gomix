package styles

func Isolate() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(isolationProp): "isolation",
		}
	}
}

func IsolationAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(clearProp): "auto",
		}
	}
}
