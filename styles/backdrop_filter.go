package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func BackdropFilterNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): "none",
		}
	}
}

func BackdropFilter(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): value.Value(),
		}
	}
}

func BackdropBlurXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "xs")),
		}
	}
}
func BackdropBlurSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "sm")),
		}
	}
}
func BackdropBlurMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "md")),
		}
	}
}
func BackdropBlurLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "lg")),
		}
	}
}
func BackdropBlurXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "xl")),
		}
	}
}
func BackdropBlur2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "2xl")),
		}
	}
}
func BackdropBlur3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", s.theme.UseVarKey(themes.Blur, "3xl")),
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

func BackdropBlur(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(backdropFilterProp): fmt.Sprintf("blur(%s)", value.Value()),
		}
	}
}

func BackdropBrightness(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("brightness(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("brightness(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropContrast(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("contrast(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("contrast(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropGrayscale(value ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropHueRotate(value any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropInvert(value ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropSaturate(value any) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		switch v := value.(type) {
		case int, float32, float64:
			prop = fmt.Sprintf("saturate(%d%%)", v)
		case customValue:
			prop = fmt.Sprintf("saturate(%s)", v.Value())
		}

		return &properties{
			string(backdropFilterProp): prop,
		}
	}
}

func BackdropSepia(value ...any) ApplyProp {
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
			string(backdropFilterProp): prop,
		}
	}
}
