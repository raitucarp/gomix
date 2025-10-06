package styles

import "github.com/raitucarp/gomix/value"

func ObjectTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "top left",
		}
	}
}

func ObjectTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "top",
		}
	}
}

func ObjectTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "top right",
		}
	}
}

func ObjectLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "left",
		}
	}
}

func ObjectCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "center",
		}
	}
}

func ObjectRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "right",
		}
	}
}

func ObjectBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "bottom left",
		}
	}
}

func ObjectBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "bottom",
		}
	}
}

func ObjectBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): "bottom right",
		}
	}
}

func ObjectBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): val.Value(),
		}
	}
}
