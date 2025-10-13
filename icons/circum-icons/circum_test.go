package ci

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestCiBowlNoodles(t *testing.T) {
	actual := CiBowlNoodles()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "Bowl_&_Noodles")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
