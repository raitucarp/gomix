package disclosure

import (
	"embed"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/alpinejs"
)

//go:embed scripts/*
var disclosureScripts embed.FS

func UseDisclosureComponents() gomix.AppParam {
	return func(app *gomix.Application) (scope gomix.Scope, fn func(params ...any)) {
		return gomix.WebScope, func(params ...any) {

			app.Apply(scope, alpinejs.GlobalsEmbed(disclosureScripts))
		}
	}
}
