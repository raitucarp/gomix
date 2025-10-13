package bi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestBiSolidAlarmAdd(t *testing.T) {
	actual := BiSolidAlarmAdd()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "solid", "bxs-alarm-add")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestBiZoomOut(t *testing.T) {
	actual := BiZoomOut()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "regular", "bx-zoom-out")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
