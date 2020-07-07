// Package gui provides ...
package gui

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

const menu_item_format = "%d %s\n"

const side_width = 30
const status_height = 2

func DrawSideGUi(g *gocui.Gui, config *model.AppConfig, isInit bool) error {
	_, maxY := g.Size()
	if isInit {
		v, err := g.SetView(model.SIDE_VIEW, 0, 0, side_width, maxY-status_height-1)
		if err != nil {
			drawSide(g, v, config)
			if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
				return err
			}
		}
	} else {
		v, err := g.View(model.SIDE_VIEW)
		if err == nil {
			drawSide(g, v, config)
		}
	}
	return nil

}
func drawSide(g *gocui.Gui, v *gocui.View, config *model.AppConfig) error {
	v.Clear()
	v.Title = "Process list"
	v.Highlight = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
	if len(config.ListProcess) == 0 {
		fmt.Fprintln(v, "No process")
	}
	for i, c := range config.ListProcess {
		fmt.Fprintf(v, menu_item_format, i+1, c.Name)
	}

	return nil
}

func DrawMainGui(g *gocui.Gui, config *model.AppConfig, isInit bool) error {
	maxX, maxY := g.Size()
	if isInit {
		v, err := g.SetView(model.MAIN_VIEW, side_width+1, 0, maxX-1, maxY/4-1)
		if err != nil {
			drawMain(g, v, config)
		}
	} else {
		v, err := g.View(model.MAIN_VIEW)
		if err == nil {
			drawMain(g, v, config)
		}
	}
	return nil
}

func drawMain(g *gocui.Gui, v *gocui.View, config *model.AppConfig) error {
	v.Clear()
	v.Title = "Main"
	status := ColoredStringDirect(config.Status.String(), color.New(color.FgHiGreen, color.BgRed))
	fmt.Fprintf(v, "Status : %s ", status)
	if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
		return err
	}
	return nil
}

func DrawStatusGui(g *gocui.Gui, config *model.AppConfig, isInit bool) error {
	maxX, maxY := g.Size()
	if isInit {
		v, err := g.SetView(model.STATUS_VIEW, -1, maxY-status_height, maxX, maxY)
		if err != nil && v != nil {
			drawStatus(g, v, config)
		}
	} else {
		v, err := g.View(model.STATUS_VIEW)
		if err == nil {
			drawStatus(g, v, config)
		}
	}
	return nil
}

func drawStatus(g *gocui.Gui, v *gocui.View, config *model.AppConfig) error {
	v.Clear()
    helpMessgage:="Press H for Help."
	status := ColoredStringDirect(config.Status.String(), color.New(color.FgHiGreen, color.BgRed))
	fmt.Fprintf(v, "Status: %s Message: %s         %s", status, config.Message,helpMessgage)
	return nil
}

func DrawProcessGui(g *gocui.Gui, config *model.AppConfig, isInit bool) error {
	maxX, maxY := g.Size()
	if isInit {
		v, err := g.SetView(model.PROCESS_VIEW, side_width+1, 0, maxX-1, maxY/2-1)
		if err != nil {
			drawProcess(g, v, config)
		}
	} else {
		v, err := g.View(model.PROCESS_VIEW)
		if err == nil {
			drawProcess(g, v, config)
		}

	}
	return nil
}

func drawProcess(g *gocui.Gui, v *gocui.View, config *model.AppConfig) error {
	v.Clear()
	if config.SelectedProcess != nil {
		v.Title = "Process "
		v.Wrap = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintf(v, "PID   : %d \n", config.SelectedProcess.Pid)
		fmt.Fprintf(v, "WID   : %s \n", config.SelectedProcess.Wid)
		fmt.Fprintf(v, "Name  : %s \n", SubString(config.SelectedProcess.Name, 50))
		fmt.Fprintf(v, "Title : %s \n", SubString(config.SelectedProcess.Title, 50))
		fmt.Fprintf(v, "Time  : %d \n", config.SelectedProcess.Time)
		fmt.Fprint(v, "Points  : ")
		for _, p := range config.SelectedProcess.Points {
			fmt.Fprintf(v, "[%d %d] ,", p.X, p.Y)
		}
	}
	return nil
}

func DrawEditorGui(g *gocui.Gui, config *model.AppConfig, isInit bool) error {
	maxX, maxY := g.Size()
	if isInit {
		v, err := g.SetView(model.EDITOR_VIEW, side_width+1, maxY/2, maxX-1, maxY-status_height-1)
		if err != nil {
			drawEditor(g, v, config)
		}
	} else {
		v, err := g.View(model.EDITOR_VIEW)
		if err == nil {
			drawEditor(g, v, config)
		}
	}
	return nil
}


func drawEditor(g *gocui.Gui, v *gocui.View, config *model.AppConfig) error {
	v.Clear()
	v.Editable = true
	v.Wrap = true
	if config.SelectedProcess != nil {
		v.Title = "Editor"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprint(v, config.SelectedProcess.Text)
		if len(config.SelectedProcess.Text) == 0 {
			v.SetCursor(0, 0)
		}
	}
	return nil
}

func DrawSideGui(g *gocui.Gui, config *model.AppConfig) error {
	if v, err := g.View(model.SIDE_VIEW); err != nil {
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
