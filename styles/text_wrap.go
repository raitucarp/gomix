package styles

func TextWrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textWrapProp): "wrap",
		}
	}
}

func TextNoWrap() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textWrapProp): "no-wrap",
		}
	}
}

func TextBalance() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textWrapProp): "balance",
		}
	}
}

func TextPretty() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textWrapProp): "pretty",
		}
	}
}
