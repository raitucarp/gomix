package value

import "testing"

func TestMul(t *testing.T) {
	actual := Mul(Number(1), Number(2)).Value()
	expected := "1 * 2"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
func TestFraction(t *testing.T) {
	actual := Fraction(Number(1), Number(2)).Value()
	expected := "1 / 2"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}

func TestPlus(t *testing.T) {
	actual := Add(Number(1), Number(2)).Value()
	expected := "1 + 2"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
func TestMinus(t *testing.T) {
	actual := Subtract(Number(1), Number(2)).Value()
	expected := "1 - 2"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}

func TestNeg(t *testing.T) {
	actual := Neg(Number(2)).Value()
	expected := "-2"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
