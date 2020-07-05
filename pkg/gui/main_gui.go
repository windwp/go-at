package gui

import (
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

func SetUpGui(g *gocui.Gui, config *model.AppConfig) error {
	DrawSideGUi(g, config, true)
	DrawMainGui(g, config, true)
	DrawProcessGui(g, config, true)
	DrawEditorGui(g, config, true)
	return nil
}

func NextView(g *gocui.Gui, v *gocui.View) error {
	cV := ""
	if v == nil {
		cV = model.SIDE_VIEW
	} else {
		cV = v.Name()
	}

	switch cV {
	case model.SIDE_VIEW:
		cV = model.MAIN_VIEW
		break
	case model.MAIN_VIEW:
		cV = model.PROCESS_VIEW
		break
	case model.PROCESS_VIEW:
		cV = model.EDITOR_VIEW
		break
	case model.EDITOR_VIEW:
		cV = model.SIDE_VIEW
		break
	default:
		cV = model.SIDE_VIEW
		break
	}
	_, err := g.SetCurrentView(cV)
	return err
}
