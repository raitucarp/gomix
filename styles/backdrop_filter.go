package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func BackdropFilterNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): "none",
		}
	}
}

func BackdropFilter(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): val.Value(),
		}
	}
}

func BackdropBlurXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "xs")),
		}
	}
}
func BackdropBlurSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "sm")),
		}
	}
}
func BackdropBlurMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "md")),
		}
	}
}
func BackdropBlurLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "lg")),
		}
	}
}
func BackdropBlurXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "xl")),
		}
	}
}
func BackdropBlur2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "2xl")),
		}
	}
}
func BackdropBlur3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(theme.Blur, "3xl")),
		}
	}
}

func BackdropBlurNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): " ",
		}
	}
}

func BackdropBlur(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", val.Value()),
		}
	}
}

func BackdropBrightness(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("brightness(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("brightness(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropContrast(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("contrast(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("contrast(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropGrayscale(val ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropHueRotate(val any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropInvert(val ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropSaturate(val any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := val.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("saturate(%d%%)", v)
		case value.Value:
			prop = fmt.Sprintf("saturate(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropSepia(val ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}
