package ri

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestRiWirelessChargingLine(t *testing.T) {
	actual := RiWirelessChargingLine()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "wireless-charging-line")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
