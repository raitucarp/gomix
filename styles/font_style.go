package styles

func Italic() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStyleProp): "italic",
		}
	}
}

func NotItalic() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStyleProp): "normal",
		}
	}
}
