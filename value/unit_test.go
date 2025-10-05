package value

import "testing"

func TestRem(t *testing.T) {
	actual := Rem(2.5).Value()
	expected := "2.5rem"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
func TestPx(t *testing.T) {
	actual := Px(96).Value()
	expected := "96px"

	if actual != expected {
		t.Errorf("actual %s is not equal %s", actual, expected)
	}
}
