package styles

import "github.com/raitucarp/gomix/value"

func WillChangeAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): "auto",
		}
	}
}

func WillChangeScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): "scroll-position",
		}
	}
}

func WillChangeContents() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): "contents",
		}
	}
}

func WillChangeTransform() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): "transform",
		}
	}
}

func WillChange(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): val.Value(),
		}
	}
}
