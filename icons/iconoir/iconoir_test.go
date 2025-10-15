package ico

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestIcoWindowNoAccess(t *testing.T) {
	actual := IcoWindowNoAccess()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, regularIcon, "window-no-access")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestIcoSolidWarningTriangle(t *testing.T) {
	actual := IcoSolidWarningTriangle()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, solidIcon, "warning-triangle")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
