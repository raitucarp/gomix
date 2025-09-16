package addons

import (
	"github.com/raitucarp/gomix"
)

func Htmx(app *gomix.App) {
	app.Scripts("https://cdn.jsdelivr.net/npm/htmx.org@2.0.7/dist/htmx.min.js")
}
