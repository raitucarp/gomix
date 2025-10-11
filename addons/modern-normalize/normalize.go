package modern_normalize

import (
	"github.com/raitucarp/gomix"
)

const AddonName = "modern-normalize"
const cdnUrl = "https://cdn.jsdelivr.net/npm/modern-normalize/modern-normalize.min.css"

func Addon() gomix.AppParam {
	return func(app *gomix.Application) (scope gomix.Scope, fn func(params ...any)) {
		return gomix.WebScope, func(params ...any) {
			app.InstallAddon(AddonName)

			app.Apply(scope, gomix.Stylesheets(cdnUrl))
		}
	}
}
