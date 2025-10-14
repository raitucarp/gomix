package lia

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestLiaOmSolid(t *testing.T) {
	actual := LiaOmSolid()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "om-solid")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
