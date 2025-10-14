package cg

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "CSS.gg Icons"
)

type cssGGIcon struct {
	el *element.HtmlElement
}

func (i *cssGGIcon) IsIcon()                       {}
func (i *cssGGIcon) Name() string                  { return name }
func (i *cssGGIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &cssGGIcon{el}
	return icon
}
