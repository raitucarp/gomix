package rx

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath     = "svg"
	name        = "Radix Icons"
)

type radixIcon struct {
	el *element.HtmlElement
}

func (i *radixIcon) IsIcon()                       {}
func (i *radixIcon) Name() string                  { return name }
func (i *radixIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &radixIcon{el}
	return icon
}
