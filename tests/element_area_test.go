package gomix_test

import (
	"testing"

	"github.com/raitucarp/gomix/element"
)

func TestElement_AreaSimple(t *testing.T) {
	actual := element.Area().Render()
	expected := "<area/>"

	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestElementArea_GlobalAttributes(t *testing.T) {
	area := element.Area()

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"id attribute",
			element.Element(area).Id("test").Render(),
			`<area id="test"/>`},
		{"class attribute",
			element.Element(area).Class("test", "test2").Render(),
			`<area id="test" class="test test2"/>`},
		{"data-name attribute",
			element.Element(area).Data("name", "test").Render(),
			`<area id="test" class="test test2" data-name="test"/>`},
		{"data-dummy attribute",
			element.Element(area).Data("dummy", "test").Render(),
			`<area id="test" class="test test2" data-name="test" data-dummy="test"/>`},
		{"dir attribute",
			element.Element(area).Dir(element.DirectionLtr).Render(),
			`<area id="test" class="test test2" data-name="test" data-dummy="test" dir="ltr"/>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, tt.actual)
			}
		})
	}
}
