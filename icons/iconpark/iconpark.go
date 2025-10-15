package icp

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath     = "svg"
	name        = "IconPark"
)

type iconPark struct {
	el *element.HtmlElement
}

func (i *iconPark) IsIcon()                       {}
func (i *iconPark) Name() string                  { return name }
func (i *iconPark) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &iconPark{el}
	return icon
}
