package styles

import "fmt"

func AutoRowsAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): "auto",
		}
	}
}

func AutoRowsMin() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): "min-content",
		}
	}
}

func AutoRowsMax() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): "max-content",
		}
	}
}

func AutoRowsFr() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): "min-max(0,1fr)",
		}
	}
}

func AutoRowsCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func AutoRowsValue(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(gridAutoRowsProp): value.String(),
		}
	}
}
