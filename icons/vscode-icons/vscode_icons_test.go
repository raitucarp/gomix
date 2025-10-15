package vsc

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestVscActivateBreakpoints(t *testing.T) {
	actual := VscActivateBreakpoints()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "activate-breakpoints")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
