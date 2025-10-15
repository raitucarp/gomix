package li

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestLiCertificateBadge1(t *testing.T) {
	actual := LiCertificateBadge1()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "certificate-badge-1")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
