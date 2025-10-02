package styles

func SelectNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(userSelectProp): "none",
		}
	}
}

func SelectText() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(userSelectProp): "text",
		}
	}
}

func SelectAll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(userSelectProp): "all",
		}
	}
}

func SelectAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(userSelectProp): "auto",
		}
	}
}
