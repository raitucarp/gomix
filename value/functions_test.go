package value

import "testing"

func TestCalc(t *testing.T) {
	t.Run("% and px", func(t *testing.T) {
		actual := Calc(
			Subtract(
				Percent(100),
				Px(80),
			),
		).Value()
		expected := "calc(100% - 80px)"

		if actual != expected {
			t.Errorf("actual %s is not equal %s", actual, expected)
		}
	})

	t.Run("var and number", func(t *testing.T) {
		actual := Calc(
			Add(
				Var("hue"),
				Number(180),
			),
		).Value()

		expected := "calc(var(--hue) + 180)"

		if actual != expected {
			t.Errorf("actual %s is not equal %s", actual, expected)
		}
	})

}
