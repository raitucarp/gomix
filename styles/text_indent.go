package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Indent(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textIndentProp): fmt.Sprintf("calc(var(--spacing) * %d)", number),
		}
	}
}

func NegIndent(number int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textIndentProp): fmt.Sprintf("calc(var(--spacing) * -%d)", number),
		}
	}
}

func IndentPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textIndentProp): "1px",
		}
	}
}

func NegIndentPx() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textIndentProp): "-1px",
		}
	}
}

func IndentBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textIndentProp): val.Value(),
		}
	}
}
