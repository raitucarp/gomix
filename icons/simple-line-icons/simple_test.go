package sl

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestSlBasketLoaded(t *testing.T) {
	actual := SlBasketLoaded()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "basket-loaded")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
