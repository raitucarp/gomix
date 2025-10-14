package fi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestFiVideoOff(t *testing.T) {
	actual := FiVideoOff()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "video-off")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
