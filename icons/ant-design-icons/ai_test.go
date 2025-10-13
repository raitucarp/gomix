package ai

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestAiFillAccountbook(t *testing.T) {
	actual := AiFillAccountBook()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, filledIcon, "account-book")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
