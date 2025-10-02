package styles

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

func WillChange(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(willChangeProp): value.Value(),
		}
	}
}
