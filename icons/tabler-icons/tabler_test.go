package tb

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestTbZoomExclamationFilled(t *testing.T) {
	actual := TbZoomExclamationFilled()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, filledIcons, "zoom-exclamation")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestTbAdjustmentsHorizontal(t *testing.T) {
	actual := TbAdjustmentsHorizontal()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, outlineIcons, "adjustments-horizontal")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
