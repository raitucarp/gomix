package wi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestWiWindBeaufort0(t *testing.T) {
	actual := WiWindBeaufort0()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "wi-wind-beaufort-0")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
