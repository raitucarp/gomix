package cg

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestCgFormatIndentIncrease(t *testing.T) {
	actual := CgFormatIndentIncrease()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "format-indent-increase")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
