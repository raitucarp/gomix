package ci

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "Circum Icons"
)

type circumIcon struct {
	el *element.HtmlElement
}

func (i *circumIcon) IsIcon()                       {}
func (i *circumIcon) Name() string                  { return name }
func (i *circumIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &circumIcon{el}
	return icon
}
