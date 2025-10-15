package sl

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svgs/*
var svgFiles embed.FS

const (
	svgPath = "svgs"
	name    = "Simple Line Icons"
)

type simpleLineIcon struct {
	el *element.HtmlElement
}

func (i *simpleLineIcon) IsIcon()                       {}
func (i *simpleLineIcon) Name() string                  { return name }
func (i *simpleLineIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &simpleLineIcon{el}
	return icon
}
