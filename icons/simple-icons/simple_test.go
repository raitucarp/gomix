package si

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestSiAbdownloadmanager(t *testing.T) {
	actual := SiAbdownloadmanager()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "abdownloadmanager")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
