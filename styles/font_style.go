package styles

func Italic() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontStyleProp): "italic",
		}
	}
}

func NotItalic() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontStyleProp): "normal",
		}
	}
}
