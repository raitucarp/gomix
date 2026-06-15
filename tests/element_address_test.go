package gomix_test

import (
	"testing"

	"github.com/raitucarp/gomix/element"
)

func TestElement_AddressSimple(t *testing.T) {
	actual := element.Address(element.Text("address")).Render()
	expected := "<address>address</address>"

	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestElementAddress_GlobalAttributes(t *testing.T) {

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"id attribute",
			element.Element(element.Address(element.Text("test"))).Id("test").Render(),
			`<address id="test">test</address>`},
		{"class attribute",
			element.Element(element.Address(element.Text("test"))).Class("test", "test2").Render(),
			`<address class="test test2">test</address>`},
		{"data-name attribute",
			element.Element(element.Address(element.Text("test"))).Data("name", "test").Render(),
			`<address data-name="test">test</address>`},
		{"data-dummy attribute",
			element.Element(element.Address(element.Text("test"))).Data("dummy", "test").Render(),
			`<address data-dummy="test">test</address>`},
		{"dir attribute",
			element.Element(element.Address(element.Text("test"))).Dir(element.DirectionLtr).Render(),
			`<address dir="ltr">test</address>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, tt.actual)
			}
		})
	}
}
