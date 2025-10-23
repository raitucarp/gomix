package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func Mask(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskImageProp): val.Value(),
		}
	}
}

func MaskNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskImageProp): "none",
		}
	}
}

func MaskLinear(degree int) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string
		if degree >= 0 {
			propValue = fmt.Sprintf("linear-gradient(%ddeg, black %s, transparent %s)", degree,
				s.theme.UseVarKey(themeCustom, "mask-linear-from"),
				s.theme.UseVarKey(themeCustom, "mask-linear-to"),
			)
		} else {
			propValue = fmt.Sprintf("linear-gradient(calc(%ddeg * -1), black %s, transparent %s)", degree,
				s.theme.UseVarKey(themeCustom, "mask-linear-from"),
				s.theme.UseVarKey(themeCustom, "mask-linear-to"),
			)
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLinearFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(%s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themeCustom, "mask-linear-position"), v,
				s.theme.UseVarKey(themeCustom, "mask-linear-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-linear-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(%s, %s %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-linear-from"),
					s.theme.UseVarKey(themeCustom, "mask-linear-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-linear-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLinearTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-linear-position"),
				s.theme.UseVarKey(themeCustom, "mask-linear-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"),
					s.theme.UseVarKey(themeCustom, "mask-linear-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"),
					s.theme.UseVarKey(themeCustom, "mask-linear-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-linear-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-linear-position"),
					s.theme.UseVarKey(themeCustom, "mask-linear-from"),
					v.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskTopFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-top-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to top, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskTopTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-top-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to top, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRightFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-right-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to right, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRightTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-right-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to right, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskBottomFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to bottom, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to bottom, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskBottomTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLeftFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to left, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-left-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to left, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLeftTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-left-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to left, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskYFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black calc(var(--spacing) * %d), transparent %s), linear-gradient(to bottom, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-top-to"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to top, %s %s, transparent %s), linear-gradient(to bottom, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskYTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, %s, transparent calc(var(--spacing) * %d)), linear-gradient(to bottom, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-top-from"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, %s %s), linear-gradient(to bottom, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-top-to"),
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s),linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-bottom-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskXFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black calc(var(--spacing) * %d), transparent %s), linear-gradient(to left, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themeCustom, "mask-right-to"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-left-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to right, %s %s, transparent %s), linear-gradient(to left, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskXTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent calc(var(--spacing) * %d)), linear-gradient(to left, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-right-from"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-left-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, %s %s), linear-gradient(to left, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-right-to"),
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s),linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-left-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskRadial(val ...any) ApplyProp {
	return func(s *Style) StyleProp {

		propValue := ""
		propKey := maskImageProp
		if len(val) >= 2 {
			propKey = maskRadialSizeProp
			propValue = fmt.Sprintf("%s %s", val[0].(value.Value).Value(), val[1].(value.Value).Value())
		} else {
			switch v := val[0].(type) {
			case int, float32, float64:
				propKey = maskRadialSizeProp
				propValue = v.(value.Value).Value()
			case *value.LiteralValue:
				propValue = v.Value()
			case *value.CustomProperty:
				propValue = v.Value()
			}
		}

		return &Properties{
			string(propKey): propValue,
		}
	}
}

func MaskRadialFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
				s.theme.UseVarKey(themeCustom, "mask-radial-size"),
				s.theme.UseVarKey(themeCustom, "mask-radial-position"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-radial-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-radial-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, %s %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-radial-from"),
					s.theme.UseVarKey(themeCustom, "mask-radial-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s), transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-radial-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRadialTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
				s.theme.UseVarKey(themeCustom, "mask-radial-size"),
				s.theme.UseVarKey(themeCustom, "mask-radial-position"),
				s.theme.UseVarKey(themeCustom, "mask-radial-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					s.theme.UseVarKey(themeCustom, "mask-radial-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, %s %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					s.theme.UseVarKey(themeCustom, "mask-radial-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-radial-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-radial-shape"),
					s.theme.UseVarKey(themeCustom, "mask-radial-size"),
					s.theme.UseVarKey(themeCustom, "mask-radial-position"),
					s.theme.UseVarKey(themeCustom, "mask-radial-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskCircle() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialShapeProp): "circle",
		}
	}
}
func MaskEllipse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialShapeProp): "ellipse",
		}
	}
}
func MaskRadialClosestCorner() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialSizeProp): "closest-corner",
		}
	}
}
func MaskRadialClosestSide() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialSizeProp): "closest-side",
		}
	}
}
func MaskRadialFarthestCorner() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialSizeProp): "farthest-corner",
		}
	}
}
func MaskRadialFarthestSide() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialSizeProp): "farthest-side",
		}
	}
}
func MaskRadialAtTopLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "top left",
		}
	}
}
func MaskRadialAtTop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "top",
		}
	}
}
func MaskRadialAtTopRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "top right",
		}
	}
}
func MaskRadialAtLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "left",
		}
	}
}
func MaskRadialAtCenter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "center",
		}
	}
}
func MaskRadialAtRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "right",
		}
	}
}
func MaskRadialAtBottomLeft() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "bottom left",
		}
	}
}
func MaskRadialAtBottom() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "bottom",
		}
	}
}
func MaskRadialAtBottomRight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(maskRadialPositionProp): "bottom right",
		}
	}
}

func MaskConic(number float64) ApplyProp {
	return func(s *Style) StyleProp {
		var prop string
		if number < 0 {
			prop = fmt.Sprintf("conic-gradient(from calc(%#vdeg * -1), black %s, transparent %s)",
				number,
				s.theme.UseVarKey(themeCustom, "mask-conic-from"),
				s.theme.UseVarKey(themeCustom, "mask-conic-to"),
			)
		} else {
			prop = fmt.Sprintf("conic-gradient(from %#vdeg, black %s, transparent %s)",
				number,
				s.theme.UseVarKey(themeCustom, "mask-conic-from"),
				s.theme.UseVarKey(themeCustom, "mask-conic-to"),
			)

		}
		return &Properties{
			string(maskImageProp): prop,
		}
	}
}

//--

func MaskConicFrom(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("conic-gradient(from %s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themeCustom, "mask-conic-position"),
				v,
				s.theme.UseVarKey(themeCustom, "mask-conic-to"),
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-conic-to"),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("conic-gradient(from %s, %s %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-conic-from"),
					s.theme.UseVarKey(themeCustom, "mask-conic-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-conic-to"),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskConicTo(val any) ApplyProp {
	return func(s *Style) StyleProp {
		var propValue string

		switch v := val.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themeCustom, "mask-conic-position"),
				s.theme.UseVarKey(themeCustom, "mask-conic-from"),
				v,
			)

		case value.Value:
			switch vv := v.(type) {
			case *value.PercentageValue:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					s.theme.UseVarKey(themeCustom, "mask-conic-from"),
					vv.Value(),
				)
			case *value.RgbaColor, *value.HslaColor, *value.Hexcolor:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, %s %s",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					s.theme.UseVarKey(themeCustom, "mask-conic-from"),
					vv.Value(),
					s.theme.UseVarKey(themeCustom, "mask-conic-to"),
				)
			case *value.CustomProperty, *value.LiteralValue:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themeCustom, "mask-conic-position"),
					s.theme.UseVarKey(themeCustom, "mask-conic-from"),
					vv.Value(),
				)
			}
		}
		return &Properties{
			string(maskImageProp): propValue,
		}
	}
}
