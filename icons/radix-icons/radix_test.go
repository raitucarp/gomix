package rx

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestRxAlignHorizontalCenters(t *testing.T) {
	actual := RxAlignHorizontalCenters()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "align-horizontal-centers")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
