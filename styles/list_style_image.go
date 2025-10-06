package styles

import "github.com/raitucarp/gomix/value"

func ListImage(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(listStyleImageProp): val.Value(),
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
