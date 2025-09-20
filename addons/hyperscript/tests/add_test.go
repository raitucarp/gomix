package hyperscript_test

import (
	"testing"

	h "github.com/raitucarp/gomix/addons/hyperscript"
)

func TestOnClickAddClass(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).Command(
			h.Add(h.Class("clicked")),
		),
	)

	expected := "on click add .clicked"

	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestOnClickAddClassToId(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).Command(
			h.Add(h.Class("clicked")).
				To(h.Id("another-div")),
		),
	)

	expected := "on click add .clicked to #another-div"

	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestOnClickAddDisabled(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).Command(
			h.Add(h.Attribute().Disabled()),
		),
	)

	expected := "on click add @disabled"

	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}

func TestOnClickAccentColorToBody(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).Command(
			h.Add(
				h.CSSObject("--accent-color", h.AttributeOf(h.My()).Value()),
			).To(h.My().Attribute()),
		),
	)

	expected := "on change add { --accent-color: my.value } to document.body"

	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}
