package gomix_test

import (
	"testing"

	"github.com/raitucarp/gomix/element"
)

func TestElementA_Simple(t *testing.T) {
	a := element.A(element.Text("test"))
	actual := a.Render()
	expected := "<a>test</a>"

	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}

	actual = a.Children(element.Text(" testa")).Render()
	expected = "<a>test testa</a>"

	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestElementA_Attributes(t *testing.T) {
	////text := element.Text("test")

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"download attribute", element.A(element.Text("test")).Download("file.jpg").Render(), `<a download="file.jpg">test</a>`},
		{"href attribute", element.A(element.Text("test")).Href("http://example.com").Render(), `<a href="http://example.com">test</a>`},
		{"hreflang attribute", element.A(element.Text("test")).HrefLang(element.LangEnglish).Render(), `<a hreflang="en">test</a>`},
		{"media attribute", element.A(element.Text("test")).Media("print and (resolution:300dpi)").Render(), `<a media="print and (resolution:300dpi)">test</a>`},
		{"ping attribute", element.A(element.Text("test")).Ping("example.com", "test.com").Render(), `<a ping="example.com, test.com">test</a>`},
		{"referrer policy attribute", element.A(element.Text("test")).ReferrerPolicy(element.RefererPolicyNoReferer).Render(), `<a referrerpolicy="no-referrer">test</a>`},
		{"rel attribute", element.A(element.Text("test")).Rel(element.Alternate).Render(), `<a rel="alternate">test</a>`},
		{"target attribute", element.A(element.Text("test")).Target(element.TargetBlank).Render(), `<a target="_blank">test</a>`},
		{"type attribute", element.A(element.Text("test")).Type("text/html").Render(), `<a type="text/html">test</a>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, tt.actual)
			}
		})
	}
}

func TestElementA_GlobalAttributes(t *testing.T) {
	////text := element.Text("test")

	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"id attribute",
			element.Element(element.A(element.Text("test"))).Id("test").Render(),
			`<a id="test">test</a>`},
		{"class attribute",
			element.Element(element.A(element.Text("test"))).Class("test", "test2").Render(),
			`<a class="test test2">test</a>`},
		{"data-name attribute",
			element.Element(element.A(element.Text("test"))).Data("name", "test").Render(),
			`<a data-name="test">test</a>`},
		{"data-dummy attribute",
			element.Element(element.A(element.Text("test"))).Data("dummy", "test").Render(),
			`<a data-dummy="test">test</a>`},
		{"dir attribute",
			element.Element(element.A(element.Text("test"))).Dir(element.DirectionLtr).Render(),
			`<a dir="ltr">test</a>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, tt.actual)
			}
		})
	}
}
