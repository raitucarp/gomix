package gr

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestGrWheelchairActive(t *testing.T) {
	actual := GrWheelchairActive()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "wheelchair-active")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
