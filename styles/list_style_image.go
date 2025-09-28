package styles

import "fmt"

func ListImageValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): value.String(),
		}
	}
}

func ListImageCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func ListImageNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): "none",
		}
	}
}
