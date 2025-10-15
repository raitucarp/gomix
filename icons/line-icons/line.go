package li

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svgs/*
var svgFiles embed.FS

const (
	svgPath     = "svgs"
	name        = "Line Icons"
)

type lineIcon struct {
	el *element.HtmlElement
}

func (i *lineIcon) IsIcon()                       {}
func (i *lineIcon) Name() string                  { return name }
func (i *lineIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &lineIcon{el}
	return icon
}
