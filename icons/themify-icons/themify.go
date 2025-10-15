package tfi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath      = "svg"
	name         = "Themify Icons"
)

type themifyIcon struct {
	el *element.HtmlElement
}

func (i *themifyIcon) IsIcon()                       {}
func (i *themifyIcon) Name() string                  { return name }
func (i *themifyIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &themifyIcon{el}
	return icon
}
