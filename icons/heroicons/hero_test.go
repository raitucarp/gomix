package hi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestHiViewColumns(t *testing.T) {
	actual := HiViewColumns()
	svgEl := icons.CreateSvgElementByFs(svgs, svgPath, solidIcon, "view-columns")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestHiOutlineUserCircle(t *testing.T) {
	actual := HiOutlineUserCircle()
	svgEl := icons.CreateSvgElementByFs(svgs, svgPath, outlineIcon, "user-circle")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
