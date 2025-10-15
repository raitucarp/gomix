package ti

import (
	"testing"

	"github.com/raitucarp/gomix/icons"
)

func TestTiWeatherDownpour(t *testing.T) {
	actual := TiWeatherDownpour()
	svgEl := icons.CreateSvgElementByFs(svgFiles, svgPath, "weather-downpour")
	expected := svgEl

	if actual.Element().Render() != expected.Element().Render() {
		t.Errorf("%s %s", actual.Element().Render(), expected.Element().Render())
	}
}
