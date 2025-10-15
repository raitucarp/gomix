package si

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath     = "icons"
	name        = "Simple Icons"
)

type simpleIcon struct {
	el *element.HtmlElement
}

func (i *simpleIcon) IsIcon()                       {}
func (i *simpleIcon) Name() string                  { return name }
func (i *simpleIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &simpleIcon{el}
	return icon
}
