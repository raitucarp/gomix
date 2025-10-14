package di

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestCgFormatIndentIncrease(t *testing.T) {
	actual := DiSmashingMagazine()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "smashing_magazine")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
