package ev

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestEvArrowheadLeft(t *testing.T) {
	actual := EvArrowheadLeft()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, fillIcon, "arrowhead-left")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestEvArrowheadRightOutline(t *testing.T) {
	actual := EvArrowheadRightOutline()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, outlineIcon, "arrowhead-right-outline")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
