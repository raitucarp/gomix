package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func FilterNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): "none",
		}
	}
}

func Filter(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): value.Value(),
		}
	}
}

func BlurXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "xs")),
		}
	}
}
func BlurSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "sm")),
		}
	}
}
func BlurMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "md")),
		}
	}
}
func BlurLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "lg")),
		}
	}
}
func BlurXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "xl")),
		}
	}
}
func Blur2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "2xl")),
		}
	}
}
func Blur3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "3xl")),
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

func Blur(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(filterProp): fmt.Sprintf("blur(%s)", value.Value()),
		}
	}
}

func Brightness(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("brightness(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("brightness(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Contrast(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("contrast(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("contrast(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Grayscale(value ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(value) <= 0 {
			prop = "100%"
		} else {
			switch v := value[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("grayscale(%d%%)", v)
			case customValue:
				prop = fmt.Sprintf("grayscale(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func HueRotate(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
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
		case customValue:
			prop = fmt.Sprintf("hue-rotate(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Invert(value ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(value) <= 0 {
			prop = "100%"
		} else {
			switch v := value[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("invert(%d%%)", v)
			case customValue:
				prop = fmt.Sprintf("invert(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Saturate(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("saturate(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("saturate(%s)", v.Value())
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}

func Sepia(value ...any) ApplyProp {
	return func(s *style) styleProp {
		var prop string

		if len(value) <= 0 {
			prop = "100%"
		} else {
			switch v := value[0].(type) {
			case int, float32, float64:
				prop = fmt.Sprintf("sepia(%d%%)", v)
			case customValue:
				prop = fmt.Sprintf("sepia(%s)", v.Value())
			}
		}

		return &properties{
			string(filterProp): prop,
		}
	}
}
