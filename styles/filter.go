package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func FilterNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): "none",
		}
	}
}

func Filter(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): val.Value(),
		}
	}
}

func BlurXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "xs")),
		}
	}
}
func BlurSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "sm")),
		}
	}
}
func BlurMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "md")),
		}
	}
}
func BlurLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "lg")),
		}
	}
}
func BlurXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "xl")),
		}
	}
}
func Blur2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "2xl")),
		}
	}
}
func Blur3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "3xl")),
		}
	}
}

func BlurNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): " ",
		}
	}
}

func Blur(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", val.Value()),
		}
	}
}

func Brightness(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("brightness(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("brightness(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Contrast(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("contrast(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("contrast(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Grayscale(val ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(val) <= 0 {
			prop = "100%"
		} else {
			switch v := val[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("grayscale(%d%%)", v)
			case value.Value:
				prop = fmt.Sprintf("grayscale(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func HueRotate(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int:
			if v < 0 {
				prop = fmt.Sprintf("hue-rotate(calc(%ddeg * -1))", v)
			} else {
				prop = fmt.Sprintf("hue-rotate(%ddeg)", v)
			}
		case float32:
			if v < 0 {
				prop = fmt.Sprintf("hue-rotate(calc(%fdeg * -1))", v)
			} else {
				prop = fmt.Sprintf("hue-rotate(%fdeg)", v)
			}
		case float64:
			if v < 0 {
				prop = fmt.Sprintf("hue-rotate(calc(%fdeg * -1))", v)
			} else {
				prop = fmt.Sprintf("hue-rotate(%fdeg)", v)
			}
		case value.Value:
			prop = fmt.Sprintf("hue-rotate(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Invert(val ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(val) <= 0 {
			prop = "100%"
		} else {
			switch v := val[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("invert(%d%%)", v)
			case value.Value:
				prop = fmt.Sprintf("invert(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Saturate(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("saturate(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("saturate(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Sepia(val ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(val) <= 0 {
			prop = "100%"
		} else {
			switch v := val[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("sepia(%d%%)", v)
			case value.Value:
				prop = fmt.Sprintf("sepia(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}
