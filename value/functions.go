package value

import (
	"fmt"
	"strings"
)

type function struct {
	name   string
	params []Value
	format string
}

func createFunction(name string, format string, params ...Value) *function {
	return &function{name: name, format: format, params: params}
}

func valuesToCommaDelimitedFormat(params ...Value) string {
	f := strings.Repeat("%s, ", len(params))
	f = strings.Trim(f, ", ")

	return "%s(" + f + ")"
}

func valuesToSpaceDelimitedFormat(params ...Value) string {
	f := strings.Repeat("%s ", len(params))
	f = strings.Trim(f, " ")

	return "%s(" + f + ")"
}

func spreadNameValue(name string, values ...Value) []any {
	params := []any{name}
	for _, v := range values {
		params = append(params, v.Value())
	}

	return params
}

func (fn *function) Value() string {
	return fmt.Sprintf(fn.format, spreadNameValue(fn.name, fn.params...)...)
}

func Abs(calcSum Value) *function {
	return createFunction("abs", "%s(%s)", calcSum)
}

func Acos(number Value) *function {
	return createFunction("acos", "%s(%s)", number)
}

func AnchorSize(params ...Value) *function {
	format := "%s(%s)"
	switch len(params) {
	case 2:
		format = "%s(%s %s)"
	case 3:
		format = "%s(%s %s, %s)"
	}
	return createFunction("anchor-size", format, params...)
}

func Anchor(params ...Value) *function {
	format := "%s(%s)"
	switch len(params) {
	case 2:
		format = "%s(%s %s)"
	case 3:
		format = "%s(%s %s, %s)"
	}
	return createFunction("anchor", format, params...)
}

func Asin(number Value) *function {
	return createFunction("anchor-size", "%s(%s)", number)
}

func Atan(number Value) *function {
	return createFunction("atan", "%s(%s)", number)
}

func Atan2(y Value, x Value) *function {
	return createFunction("atan2", "%s(%s, %s)", y, x)
}

func Attr(name Value, params ...Value) *function {
	format := "%s(%s)"
	switch len(params) {
	case 1:
		switch params[0].(type) {
		case typ:
			format = "%s(%s %s)"
		default:
			format = "%s(%s, %s)"
		}
	case 2:
		format = "%s(%s %s, %s)"
	}

	return createFunction("attr", format, params...)
}

func Blur(length Value) *function {
	return createFunction("blur", "%s(%s)", length)
}

func Brightness(length Value) *function {
	return createFunction("brightness", "%s(%s)", length)
}

func Calc(param Value) *function {
	return createFunction("calc", "%s(%s)", param)
}

func Circle(param Value) *function {
	return createFunction("circle", "%s(%s)", param)
}

func Clamp(min Value, val Value, max Value) *function {
	return createFunction("clamp", "%s(%s, %s, %s)", min, val, max)
}

func ColorMix(params ...Value) *function {

	return createFunction("color-mix", valuesToCommaDelimitedFormat(params...), params...)
}

func Color(params ...Value) *function {
	return createFunction("color", valuesToSpaceDelimitedFormat(params...), params...)
}

func ConicGradient(params ...Value) *function {
	return createFunction("conic-gradient", valuesToCommaDelimitedFormat(params...), params...)
}

func Contrast(number Value) *function {
	return createFunction("contrast", "%s(%s)", number)
}

func Cos(angle Value) *function {
	return createFunction("cos", "%s(%s)", angle)
}

func CubicBezier(x1 Value, y1 Value, x2 Value, y2 Value) *function {
	return createFunction("cubic-bezier", "%s(%s, %s, %s, %s)", x1, x2, y1, y2)
}
