package styles

func TextWrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textWrapProp): "wrap",
		}
	}
}

func TextNoWrap() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textWrapProp): "no-wrap",
		}
	}
}

func TextBalance() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textWrapProp): "balance",
		}
	}
}

func TextPretty() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textWrapProp): "pretty",
		}
	}
}
