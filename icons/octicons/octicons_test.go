package octicons

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestGoFileDirectoryOpenFill(t *testing.T) {
	actual := GoFileDirectoryOpenFill()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "file-directory-open-fill")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
