package fa6

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestFaBrand500px(t *testing.T) {
	actual := FaBrand500px()
	svgEl := icons.CreateSvgElementByFs(svgs, svgPath, "brands", "500px")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
