package octicons

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "Github Octicons icons"
)

type octicons struct {
	el *element.HtmlElement
}

func (i *octicons) IsIcon()                       {}
func (i *octicons) Name() string                  { return name }
func (i *octicons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &octicons{el}
	return icon
}
