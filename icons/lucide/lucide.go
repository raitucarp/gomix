package lu

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "Lucide Icons"
)

type lucideIcon struct {
	el *element.HtmlElement
}

func (i *lucideIcon) IsIcon()                       {}
func (i *lucideIcon) Name() string                  { return name }
func (i *lucideIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &lucideIcon{el}
	return icon
}
