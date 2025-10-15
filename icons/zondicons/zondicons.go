package zi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "Zondicons"
)

type zondicons struct {
	el *element.HtmlElement
}

func (i *zondicons) IsIcon()                       {}
func (i *zondicons) Name() string                  { return name }
func (i *zondicons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &zondicons{el}
	return icon
}
