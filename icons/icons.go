package icons

import (
	"embed"
	"path"

	"github.com/raitucarp/gomix/element"
)

type IsIcon interface {
	IsIcon()
	Element() *element.HtmlElement
	Name() string
}

func readSvgIcon(fs embed.FS, svgPaths ...string) (string, error) {
	f, err := fs.ReadFile(path.Join(svgPaths...) + ".svg")
	if err != nil {
		return "", err
	}

	return string(f), err
}

func CreateSvgElementByFs(fs embed.FS, paths ...string) *element.HtmlElement {
	f, err := readSvgIcon(fs, paths...)
	if err != nil {
		panic(err)
	}
	svgEl, _ := element.SvgString(f)
	return svgEl.Element()
}
