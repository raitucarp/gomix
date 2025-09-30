package styles

import "fmt"

func propWidthValue(value ...customValue) string {
	propValue := ""
	if len(value) == 0 {
		propValue = "1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				propValue = fmt.Sprintf("%dpx", n)
			default:
				propValue = v.Value()
			}
		case *customVariableProp:
			propValue = v.Value()
		}
	}

	return propValue
}

func Border(value ...customValue) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(borderWidthProp): propWidthValue(value...),
		}
	}
}

func BorderX(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineWidthProp): propWidthValue(value...),
		}
	}
}

func BorderY(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockWidthProp): propWidthValue(value...),
		}
	}
}

func BorderS(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartWidthProp): propWidthValue(value...),
		}
	}
}

func BorderE(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndWidthProp): propWidthValue(value...),
		}
	}
}

func BorderT(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopWidthProp): propWidthValue(value...),
		}
	}
}

func BorderR(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightWidthProp): propWidthValue(value...),
		}
	}
}

func BorderB(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomWidthProp): propWidthValue(value...),
		}
	}
}

func BorderL(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftWidthProp): propWidthValue(value...),
		}
	}
}

func DivideX(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderInlineStartWidthProp): "0px",
			notLastChildProp(borderInlineEndWidthProp):   propWidthValue(value...),
		}
	}
}

func DivideY(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderTopWidthProp):    "0px",
			notLastChildProp(borderBottomWidthProp): propWidthValue(value...),
		}
	}
}

func DivideXReverse(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(divideXReverseProp): "1",
		}
	}
}

func DivideYReverse(value ...customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(divideYReverseProp): "1",
		}
	}
}
