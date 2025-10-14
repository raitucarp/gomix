package gi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestGiAbstract014(t *testing.T) {
	actual := GiAbstract014()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "abstract-014")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
