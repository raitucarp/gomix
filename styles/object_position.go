package styles

import "fmt"

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

func ObjectCustomProperty(propertyName string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): fmt.Sprintf("var(--%s)", propertyName),
		}
	}
}

func ObjectValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(objectPositionProp): value.String(),
		}
	}
}
