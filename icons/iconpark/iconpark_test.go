package icp

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestIcpAirConditioning(t *testing.T) {
	actual := IcpAirConditioning()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "air-conditioning")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
