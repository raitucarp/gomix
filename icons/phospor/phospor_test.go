package pi

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestPiYoutubeLogoBold(t *testing.T) {
	actual := PiYoutubeLogoBold()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, boldIcon, "youtube-logo-bold")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestPiAirplaneTaxiingDuotone(t *testing.T) {
	actual := PiAirplaneTaxiingDuotone()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, duotoneIcon, "airplane-taxiing-duotone")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestPiWifiNoneLight(t *testing.T) {
	actual := PiWifiNoneLight()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, lightIcon, "wifi-none-light")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestPiAlignTopSimple(t *testing.T) {
	actual := PiAlignTopSimple()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, regularIcon, "align-top-simple")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}

func TestPiAppStoreLogoThin(t *testing.T) {
	actual := PiAppStoreLogoThin()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, thinIcon, "app-store-logo-thin")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
