package hyperscript_test

import (
	"testing"

	h "github.com/raitucarp/gomix/addons/hyperscript"
)

func TestOnClick1(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).
			Count(1),
	)

	expected := "on click 1"

	if actual != expected {
		t.Errorf("On click count 1 expected %s, but got %s", expected, actual)

	}
}

func TestOnClick2To10(t *testing.T) {
	actual := h.String(
		h.On(h.EventClick).Count(2, 10),
	)

	expected := "on click 2 to 10"

	if actual != expected {
		t.Errorf("On click count 2 to 10 expected %s, but got %s", expected, actual)

	}
}

func TestKeyUpDebouncedAtMs(t *testing.T) {
	actual := h.String(
		h.On(h.EventKeyUp).DebouncedAt(h.TimeMilliseconds(500)),
	)

	expected := "on keyup debounced at 500ms"

	if actual != expected {
		t.Errorf("%s, but got %s", expected, actual)

	}
}

func TestKeyUpDebouncedAtSeconds(t *testing.T) {
	actual := h.String(
		h.On(h.EventKeyDown).
			DebouncedAt(h.TimeSeconds(2)),
	)

	expected := "on keydown debounced at 2s"

	if actual != expected {
		t.Errorf("%s, but got %s", expected, actual)

	}
}

func TestMousemoveThrotledAtMs(t *testing.T) {
	actual := h.String(
		h.On(h.EventMouseMove).
			ThrottledAt(h.TimeMilliseconds(500)),
	)

	expected := "on mousemove throttled at 500ms"

	if actual != expected {
		t.Errorf("%s, but got %s", expected, actual)

	}
}

func TestMouseEnterAtSeconds(t *testing.T) {
	actual := h.String(
		h.On(h.EventMouseEnter).
			ThrottledAt(h.TimeSeconds(2)),
	)

	expected := "on mouseenter throttled at 2s"

	if actual != expected {
		t.Errorf("%s, but got %s", expected, actual)

	}
}
