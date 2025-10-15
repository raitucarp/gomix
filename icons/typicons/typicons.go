package ti

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath      = "svg"
	name         = "Typicons"
)

type typicons struct {
	el *element.HtmlElement
}

func (i *typicons) IsIcon()                       {}
func (i *typicons) Name() string                  { return name }
func (i *typicons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &typicons{el}
	return icon
}
