package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/themes"
)

func Mask(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskImageProp): value.Value(),
		}
	}
}

func MaskNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskImageProp): "none",
		}
	}
}

func MaskLinear(degree int) ApplyProp {
	return func(s *style) styleProp {
		var propValue string
		if degree >= 0 {
			propValue = fmt.Sprintf("linear-gradient(%ddeg, black %s, transparent %s)", degree,
				s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
				s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
			)
		} else {
			propValue = fmt.Sprintf("linear-gradient(calc(%ddeg * -1), black %s, transparent %s)", degree,
				s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
				s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
			)
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLinearFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(%s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themes.Custom, "mask-linear-position"), v,
				s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(%s, %s %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
					s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"), vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLinearTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-linear-position"),
				s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"),
					s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"),
					s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-linear-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(%s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-linear-position"),
					s.theme.UseVarKey(themes.Custom, "mask-linear-from"),
					v.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskTopFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-top-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to top, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskTopTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-top-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to top, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRightFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-right-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to right, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRightTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-right-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to right, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskBottomFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to bottom, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to bottom, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskBottomTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLeftFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to left, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-left-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to left, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskLeftTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-left-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to left, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskYFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, black calc(var(--spacing) * %d), transparent %s), linear-gradient(to bottom, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-top-to"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to top, %s %s, transparent %s), linear-gradient(to bottom, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskYTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to top, %s, transparent calc(var(--spacing) * %d)), linear-gradient(to bottom, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-top-from"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, transparent %s), linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:

				propValue = fmt.Sprintf("linear-gradient(to bottom, black %s, %s %s), linear-gradient(to bottom, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-top-to"),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to top, black %s, transparent %s),linear-gradient(to bottom, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-top-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-bottom-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskXFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black calc(var(--spacing) * %d), transparent %s), linear-gradient(to left, black calc(var(--spacing) * %d), transparent %s)",
				v,
				s.theme.UseVarKey(themes.Custom, "mask-right-to"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-left-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to right, %s %s, transparent %s), linear-gradient(to left, %s %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskXTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent calc(var(--spacing) * %d)), linear-gradient(to left, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-right-from"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-left-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, transparent %s), linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("linear-gradient(to left, black %s, %s %s), linear-gradient(to left, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-right-to"),
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("linear-gradient(to right, black %s, transparent %s),linear-gradient(to left, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-right-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-left-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp):     propValue,
			string(maskCompositeProp): "intersect",
		}
	}
}

func MaskRadial(value ...any) ApplyProp {
	return func(s *style) styleProp {

		propValue := ""
		propKey := maskImageProp
		if len(value) >= 2 {
			propKey = maskRadialSizeProp
			propValue = fmt.Sprintf("%s %s", value[0].(customValue).Value(), value[1].(customValue).Value())
		} else {
			switch v := value[0].(type) {
			case *percentage, int, float32, float64:
				propKey = maskRadialSizeProp
				propValue = v.(customValue).Value()
			case *val:
				propValue = v.Value()
			}
		}

		return &properties{
			string(propKey): propValue,
		}
	}
}

func MaskRadialFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
				s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
				s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-radial-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-radial-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, %s %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-radial-from"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s), transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-radial-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskRadialTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
				s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
				s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
				s.theme.UseVarKey(themes.Custom, "mask-radial-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, %s %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-radial-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("radial-gradient(%s %s at %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-radial-shape"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-size"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-position"),
					s.theme.UseVarKey(themes.Custom, "mask-radial-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskCircle() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialShapeProp): "circle",
		}
	}
}
func MaskEllipse() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialShapeProp): "ellipse",
		}
	}
}
func MaskRadialClosestCorner() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialSizeProp): "closest-corner",
		}
	}
}
func MaskRadialClosestSide() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialSizeProp): "closest-side",
		}
	}
}
func MaskRadialFarthestCorner() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialSizeProp): "farthest-corner",
		}
	}
}
func MaskRadialFarthestSide() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialSizeProp): "farthest-side",
		}
	}
}
func MaskRadialAtTopLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "top left",
		}
	}
}
func MaskRadialAtTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "top",
		}
	}
}
func MaskRadialAtTopRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "top right",
		}
	}
}
func MaskRadialAtLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "left",
		}
	}
}
func MaskRadialAtCenter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "center",
		}
	}
}
func MaskRadialAtRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "right",
		}
	}
}
func MaskRadialAtBottomLeft() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "bottom left",
		}
	}
}
func MaskRadialAtBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "bottom",
		}
	}
}
func MaskRadialAtBottomRight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(maskRadialPositionProp): "bottom right",
		}
	}
}

func MaskConic(number float64) ApplyProp {
	return func(s *style) styleProp {
		var prop string
		if number < 0 {
			prop = fmt.Sprintf("conic-gradient(from calc(%ddeg * -1), black %s, transparent %s)",
				number,
				s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
				s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
			)
		} else {
			prop = fmt.Sprintf("conic-gradient(from %ddeg, black %s, transparent %s)",
				number,
				s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
				s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
			)

		}
		return &properties{
			string(maskImageProp): prop,
		}
	}
}

//--

func MaskConicFrom(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("conic-gradient(from %s, black calc(var(--spacing) * %d), transparent %s)",
				s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
				v,
				s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("conic-gradient(from %s, %s %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
					s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}

func MaskConicTo(value any) ApplyProp {
	return func(s *style) styleProp {
		var propValue string

		switch v := value.(type) {
		case int, float32, float64:
			propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent calc(var(--spacing) * %d))",
				s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
				s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
				v,
			)

		case customValue:
			switch vv := v.(type) {
			case *percentage:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
					vv.Value(),
				)
			case *rgba, *hsla, *hexcolor:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, %s %s",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
					vv.Value(),
					s.theme.UseVarKey(themes.Custom, "mask-conic-to"),
				)
			case *customVariableProp, *val:
				propValue = fmt.Sprintf("conic-gradient(from %s, black %s, transparent %s)",
					s.theme.UseVarKey(themes.Custom, "mask-conic-position"),
					s.theme.UseVarKey(themes.Custom, "mask-conic-from"),
					vv.Value(),
				)
			}
		}
		return &properties{
			string(maskImageProp): propValue,
		}
	}
}
