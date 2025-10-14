package lia

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "Icons8 Line Awesome"
)

type icons8LineAwesomeIcon struct {
	el *element.HtmlElement
}

func (i *icons8LineAwesomeIcon) IsIcon()                       {}
func (i *icons8LineAwesomeIcon) Name() string                  { return name }
func (i *icons8LineAwesomeIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &icons8LineAwesomeIcon{el}
	return icon
}
