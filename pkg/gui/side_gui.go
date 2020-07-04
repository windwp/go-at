// Package gui provides ...
package gui

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

func SetUpSideGui(g *gocui.Gui,config *model.AppConfig) error {
	_, maxY := g.Size()
    log.Printf("Config setup %d",len(config.ListProcess))
    v,_:=g.SetView(model.SIDE_VIEW, -1, -1, 30, maxY)
    v.Clear()
    log.Printf("Config length %d",len(config.ListProcess))
    v.Highlight = true
    v.SelBgColor = gocui.ColorGreen
    v.SelFgColor = gocui.ColorBlack
    for _, c := range config.ListProcess {
        fmt.Fprintln(v, c.Name)
    }
    return nil
}

func DrawSideGui(g *gocui.Gui,config *model.AppConfig) error{
    if v,err :=g.View(model.SIDE_VIEW);err !=nil{
		if err != gocui.ErrUnknownView {
			return err
		}
        v.Clear()
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		for _, c := range config.ListProcess {
			fmt.Fprintln(v, c.Name)
		}
    }

    return nil
}

func SetupListProcess(g *gocui.Gui,listItem []string) error{
    w:=40
    h:=20
	maxX, maxY := g.Size()
	if v, err := g.SetView(model.PROCESS_VIEW, maxX/2-w/2, maxY/2-h/2, maxX/2+w/2, maxY/2+h/2); err != nil {
		v.Title = "Dialog"
		if err != gocui.ErrUnknownView {
			return err
		}
        for _, l := range listItem {
            fmt.Fprintln(v,l)
        }
		if _, err := g.SetCurrentView(model.PROCESS_VIEW); err != nil {
			log.Panic("can't set view")
			return err
		}
	}

	return nil
}
