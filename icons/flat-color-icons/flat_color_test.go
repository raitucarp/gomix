package fc

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestBsWindows(t *testing.T) {
	actual := FcScatterPlot()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "scatter_plot")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
