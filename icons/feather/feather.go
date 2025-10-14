package fi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "Feather Icon"
)

type featherIcon struct {
	el *element.HtmlElement
}

func (i *featherIcon) IsIcon()                       {}
func (i *featherIcon) Name() string                  { return name }
func (i *featherIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &featherIcon{el}
	return icon
}
