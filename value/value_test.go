package value

import "testing"

func TestNumber(t *testing.T) {
	actual := Number(22.33).Value()
	expected := "22.33"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
