package gomix_test

import (
	"testing"

	"github.com/raitucarp/gomix/element"
)

func TestElement_AbbrSimple(t *testing.T) {
	actual := element.Abbr(element.Text("test")).Render()
	expected := "<abbr>test</abbr>"

	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestElementAbbr_GlobalAttributes(t *testing.T) {
	text := element.Text("test")

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"id attribute",
			element.Element(element.Abbr(text)).Id("test").Render(),
			`<abbr id="test">test</abbr>`},
		{"class attribute",
			element.Element(element.Abbr(text)).Class("test", "test2").Render(),
			`<abbr class="test test2">test</abbr>`},
		{"data-name attribute",
			element.Element(element.Abbr(text)).Data("name", "test").Render(),
			`<abbr data-name="test">test</abbr>`},
		{"data-dummy attribute",
			element.Element(element.Abbr(text)).Data("dummy", "test").Render(),
			`<abbr data-dummy="test">test</abbr>`},
		{"dir attribute",
			element.Element(element.Abbr(text)).Dir(element.DirectionLtr).Render(),
			`<abbr dir="ltr">test</abbr>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, tt.actual)
			}
		})
	}
}
