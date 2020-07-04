package app

import (
	"fmt"
	"log"

	// "log"
	"github.com/jroimartin/gocui"
	"github.com/thoas/go-funk"
	"github.com/windwp/go-at/pkg/gui"
	"github.com/windwp/go-at/pkg/model"
	"github.com/windwp/go-at/pkg/command"
)

var dialogHandler model.DialogHandler

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		_, err := g.SetCurrentView("main")
		return err
	}
	_, err := g.SetCurrentView("side")
	return err
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		n, _ := v.Line(cy + 1)
		if n != "" {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func addProcess(g *gocui.Gui, v *gocui.View) error {

    lProcess,_ :=command.GetListProcess()
    litem:=make([]string, 0)
    for _, l := range lProcess {
        litem=append(litem,l.Name)
    }
    gui.SetupListProcess(g,litem)


	return nil
}

func getSelectedItem(g *gocui.Gui) error {
	v, err := g.View("side")
	if err == nil {
		text := getSelectedText(g, v)
		for _, item := range config.ListProcess {
			if item.Name == text {
				config.SelectedProcess = &item
				break
			}
		}
		if config.SelectedProcess != nil {
			log.Print(config.SelectedProcess.Name)
		}
	}
	return nil
}

func getSelectedText(g *gocui.Gui, v *gocui.View) string {
	var l string
	var err error
	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return ""
	}
	return l

}
func deleteSeletedItem(g *gocui.Gui, v *gocui.View) error {
	getSelectedItem(g)
	if config.SelectedProcess != nil {
		config.ListProcess = funk.Filter(config.ListProcess, func(item model.ProcessConfig) bool {
			return item.Name != config.SelectedProcess.Name
		}).([]model.ProcessConfig)
		config.SelectedProcess = nil
	}
	return nil
}

func okDialog(g *gocui.Gui, v *gocui.View) error {
	if dialogHandler != nil {
		dialogHandler(g, v)
	}
	closeDialog(g, v)
	return nil
}

func closeDialog(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(model.MSG_VIEW); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
		return err
	}
	dialogHandler = nil
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func delItem(g *gocui.Gui, v *gocui.View) error {
	text := getSelectedText(g, v)
	if text != "" {
		text = fmt.Sprintf("Delete %s ?", text)
		ShowDialog(g, v, text, deleteSeletedItem)
	}
	return nil
}

func ShowDialog(
	g *gocui.Gui,
	v *gocui.View,
	message string, handler model.DialogHandler,
) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(model.MSG_VIEW, maxX/2-20, maxY/2-5, maxX/2+20, maxY/2); err != nil {
		v.Title = "Dialog"
		if err != gocui.ErrUnknownView {
			log.Panic("stupi")
			return err
		}
		fmt.Fprintln(v, message)
		fmt.Fprintln(v, " ")
		text := "[Y]es  |   [N]o"
		cursorX := 25 - len(text)
		for i := 0; i < cursorX; i++ {
			text = " " + text
		}

		v.SetCursor(cursorX+1, 2)
		fmt.Fprintf(v, "%s\n", text)
		if _, err := g.SetCurrentView(model.MSG_VIEW); err != nil {
			log.Panic("can't set view")
			return err
		}
		dialogHandler = handler
	}

	return nil
}
