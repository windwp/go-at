package gui

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

func SetUpMainGui(g *gocui.Gui,config *model.AppConfig) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(model.MAIN_VIEW, 30, -1, maxX, maxY); err != nil {
        log.Println("Setup Main GUI")
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true

        if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
            return err
        }
        if config.SelectedProcess != nil{
            fmt.Fprintln(v,config.SelectedProcess.Name)

        }else{
            fmt.Fprintf(v,"No item")
        }
	}
    return nil
}

