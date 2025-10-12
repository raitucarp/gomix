package fa5

import (
	"testing"
)

func TestFaBrand500px(t *testing.T) {
	actual := FaBrand500px()
	expected, err := readSvg("brands/500px.svg")
	if err != nil {
		t.Error(err)
	}

	if actual.Element().Render() != expected {
		t.Errorf("%s %s", actual.Element().Render(), expected)
	}
}
