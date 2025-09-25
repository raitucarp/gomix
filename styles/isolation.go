package styles

func Isolate() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(isolationProp): "isolation",
		}
	}
}

func IsolationAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(clearProp): "auto",
		}
	}
}
