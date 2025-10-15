package zi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestZiVolumeDown(t *testing.T) {
	actual := ZiVolumeDown()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "volume-down")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
