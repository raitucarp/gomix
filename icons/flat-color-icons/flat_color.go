package fc

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "Flat Color Icons"
)

type flatColorIcon struct {
	el *element.HtmlElement
}

func (i *flatColorIcon) IsIcon()                       {}
func (i *flatColorIcon) Name() string                  { return name }
func (i *flatColorIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &flatColorIcon{el}
	return icon
}
