package vsc

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "VS Code Icons"
)

type vscodeCodicons struct {
	el *element.HtmlElement
}

func (i *vscodeCodicons) IsIcon()                       {}
func (i *vscodeCodicons) Name() string                  { return name }
func (i *vscodeCodicons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &vscodeCodicons{el}
	return icon
}
