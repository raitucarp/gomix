package fl

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestFlAccessibilityFilled(t *testing.T) {
	actual := FlAccessibilityFilled()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, filledIcons, "ic_fluent_accessibility_24_filled")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestFlAddSquareMultiple(t *testing.T) {
	actual := FlAddSquareMultiple()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, regularIcons, "ic_fluent_add_square_multiple_24_regular")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
