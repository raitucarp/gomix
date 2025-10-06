package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func propWidthValue(val ...value.Value) string {
	propValue := ""
	if len(val) == 0 {
		propValue = "1px"
	} else {

		switch v := val[0].(type) {
		case *value.Unit[int]:
			propValue = fmt.Sprintf("%spx", v.Value())

		case *value.CustomProperty:
			propValue = v.Value()

		default:
			propValue = v.Value()

		}
	}

	return propValue
}

func Border(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {

		return &properties{
			string(borderWidthProp): propWidthValue(val...),
		}
	}
}

func BorderX(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineWidthProp): propWidthValue(val...),
		}
	}
}

func BorderY(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockWidthProp): propWidthValue(val...),
		}
	}
}

func BorderS(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartWidthProp): propWidthValue(val...),
		}
	}
}

func BorderE(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndWidthProp): propWidthValue(val...),
		}
	}
}

func BorderT(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopWidthProp): propWidthValue(val...),
		}
	}
}

func BorderR(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightWidthProp): propWidthValue(val...),
		}
	}
}

func BorderB(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomWidthProp): propWidthValue(val...),
		}
	}
}

func BorderL(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftWidthProp): propWidthValue(val...),
		}
	}
}

func DivideX(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderInlineStartWidthProp): "0px",
			notLastChildProp(borderInlineEndWidthProp):   propWidthValue(val...),
		}
	}
}

func DivideY(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderTopWidthProp):    "0px",
			notLastChildProp(borderBottomWidthProp): propWidthValue(val...),
		}
	}
}

func DivideXReverse(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(divideXReverseProp): "1",
		}
	}
}

func DivideYReverse(val ...value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(divideYReverseProp): "1",
		}
	}
}
