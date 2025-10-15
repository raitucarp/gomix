package tfi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestTfiViewList(t *testing.T) {
	actual := TfiViewList()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "view-list")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
