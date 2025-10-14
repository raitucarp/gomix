package gr

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed img/*
var svgFiles embed.FS

const (
	svgPath = "img"
	name    = "Grommet-Icons"
)

type grommetIcon struct {
	el *element.HtmlElement
}

func (i *grommetIcon) IsIcon()                       {}
func (i *grommetIcon) Name() string                  { return name }
func (i *grommetIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &grommetIcon{el}
	return icon
}
