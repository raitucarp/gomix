package gvi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestGviCheckSquareO(t *testing.T) {
	actual := GviCheckSquareO()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "check-square-o")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
