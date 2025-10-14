package lu

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestLuWebhookOff(t *testing.T) {
	actual := LuWebhookOff()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "webhook-off")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
