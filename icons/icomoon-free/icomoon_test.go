package im

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestImPaintFormat(t *testing.T) {
	actual := ImPaintFormat()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "013-paint-format")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
