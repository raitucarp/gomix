package alpinejs

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/styles"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

const AddonName = "Alpine.js Core"
const cdnBaseUrl = "https://cdn.jsdelivr.net/npm"
const coreUrl = cdnBaseUrl + "/alpinejs@%s/dist/cdn.min.js"

func installPlugin(name string, version string) gomix.AppParam {
	return func(app *gomix.Application) (scope gomix.Scope, fn func(params ...any)) {
		return gomix.WebScope, func(params ...any) {
			app.InstallAddon(fmt.Sprintf("Alpine.js: %s plugin", name))
			alpinejsUrl := fmt.Sprintf("%s/@alpinejs/%s@%s/dist/cdn.min.js", cdnBaseUrl, name, version)
			app.Apply(gomix.WebScope, gomix.ScriptsDefer(alpinejsUrl))
		}
	}
}

func MaskPlugin(version string) gomix.AppParam      { return installPlugin("mask", version) }
func IntersectPlugin(version string) gomix.AppParam { return installPlugin("intersect", version) }
func ResizePlugin(version string) gomix.AppParam    { return installPlugin("resize", version) }
func PersistPlugin(version string) gomix.AppParam   { return installPlugin("persist", version) }
func FocusPlugin(version string) gomix.AppParam     { return installPlugin("focus", version) }
func CollapsePlugin(version string) gomix.AppParam  { return installPlugin("collapse", version) }
func AnchorPlugin(version string) gomix.AppParam    { return installPlugin("anchor", version) }
func MorphPlugin(version string) gomix.AppParam     { return installPlugin("morph", version) }
func SortPlugin(version string) gomix.AppParam      { return installPlugin("sort", version) }

func Core(version string) gomix.AppParam {
	return func(app *gomix.Application) (scope gomix.Scope, fn func(params ...any)) {
		return gomix.WebScope, func(params ...any) {
			app.InstallAddon(AddonName)
			alpinejsUrl := fmt.Sprintf(coreUrl, version)

			app.Apply(scope, gomix.ScriptsDefer(alpinejsUrl))
			app.Apply(scope, gomix.StyleGlobal(gomix.CSS{
				"[x-cloak]": gomix.Style(styles.Hidden()),
			}))
		}
	}
}

func GlobalsEmbed(embedFiles embed.FS) gomix.AppParam {
	return func(app *gomix.Application) (scope gomix.Scope, fn func(params ...any)) {
		return gomix.WebScope, func(params ...any) {
			m := minify.New()
			m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

			initEventListener := `%s
			document.addEventListener('alpine:init', function alpineInit() {
				%s
			})`

			globalsDirectories := []string{"data", "store", "bind"}
			functions := []string{}
			initScripts := []string{}

			fs.WalkDir(embedFiles, ".", func(filePath string, dirEntry fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if dirEntry.IsDir() {
					return nil
				}

				pathSlice := strings.Split(filePath, "/")
				if len(pathSlice) < 2 {
					return nil
				}

				dirName := pathSlice[1]

				if !slices.Contains(globalsDirectories, dirName) {
					return nil
				}

				fileName := dirEntry.Name()

				fileNameWithExt := filepath.Base(fileName)
				fileExt := filepath.Ext(fileNameWithExt)
				name := strings.TrimSuffix(fileNameWithExt, fileExt)
				content, _ := embedFiles.ReadFile(filePath)
				functions = append(functions, string(content))
				initScripts = append(initScripts, fmt.Sprintf(`Alpine.%s('%s', %s)`, dirName, name, name))

				return nil
			})

			initEventListener, _ = m.String("application/javascript",
				fmt.Sprintf(initEventListener,
					strings.Join(functions, "\n"),
					strings.Join(initScripts, "\n"),
				),
			)

			// Alpine.data('dropdown', dropdown);
			app.Apply(scope, gomix.JavaScripts(initEventListener))
		}
	}
}

type alpine struct {
	el *element.HtmlElement
}

func (a *alpine) Element() *element.HtmlElement { return a.el }
func (a *alpine) Render() string                { return a.el.Render() }

func (a *alpine) Data(data any) *alpine {
	if s, ok := data.(string); ok {
		a.el.AddAttribute("x-data", s)
		return a
	}

	jsonData, _ := json.Marshal(data)
	a.el.AddAttribute("x-data", string(jsonData))
	return a
}

func (a *alpine) Init(script string) *alpine {
	a.el.AddAttribute("x-init", script)
	return a
}

func (a *alpine) Show(state string) *alpine {
	a.el.AddAttribute("x-show", state)
	return a
}

func (a *alpine) Bind(attribute string, state string) *alpine {
	a.el.AddAttribute(strings.Join([]string{"x-bind", attribute}, ":"), state)
	return a
}

func (a *alpine) On(event string, handler string) *alpine {
	a.el.AddAttribute("@"+event, handler)
	return a
}

func (a *alpine) Text(expression string) *alpine {
	a.el.AddAttribute("x-text", expression)
	return a
}

func (a *alpine) Html(html element.IsElement) *alpine {
	a.el.AddAttribute("x-html", html.Element().Render())
	return a
}

func (a *alpine) Model(model string) *alpine {
	a.el.AddAttribute("x-model", model)
	return a
}

func (a *alpine) Modelable(model string) *alpine {
	a.el.AddAttribute("x-modelable", model)
	return a
}

func (a *alpine) Effect(code string) *alpine {
	a.el.AddAttribute("x-effect", code)
	return a
}

func (a *alpine) Ignore() *alpine {
	a.el.AddAttributeBool("x-ignore")
	return a
}

func (a *alpine) Ref(name string) *alpine {
	a.el.AddAttribute("x-ref", name)
	return a
}

func (a *alpine) Cloak() *alpine {
	a.el.AddAttributeBool("x-cloak")
	return a
}

func (a *alpine) Teleport(selector string) *alpine {
	a.el.AddAttribute("x-teleport", selector)
	return a
}

func (a *alpine) If(condition string) *alpine {
	a.el.AddAttribute("x-if", condition)
	return a
}

func (a *alpine) Id(id string) *alpine {
	a.el.AddAttribute("x-id", id)
	return a
}

func Alpine(e components.IsComponent) *alpine {
	h := alpine{
		el: e.Element(),
	}
	return &h
}
