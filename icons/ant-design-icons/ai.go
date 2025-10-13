package ai

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath      = "svg"
	name         = "Ant Design Icons"
	filledIcon   = "filled"
	outlinedIcon = "outlined"
	twoToneIcon  = "twotone"
)

type antDesignIcon struct {
	el *element.HtmlElement
}

func (i *antDesignIcon) IsIcon()                       {}
func (i *antDesignIcon) Name() string                  { return name }
func (i *antDesignIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &antDesignIcon{el}
	return icon
}
