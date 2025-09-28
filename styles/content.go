package styles

import "fmt"

func Content(value string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): value,
		}
	}
}

func ContentCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ContentNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(contentProp): "none",
		}
	}
}
