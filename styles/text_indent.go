package styles

import "fmt"

func Indent(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegIndent(number int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func IndentPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): "1px",
		}
	}
}

func NegIndentPx() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): "-1px",
		}
	}
}

func IndentCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func IndentValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(textIndentProp): value.String(),
		}
	}
}
