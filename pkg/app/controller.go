package app

import (
	"fmt"
	"log"
	"regexp"

	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/command"
	"github.com/windwp/go-at/pkg/gui"
	"github.com/windwp/go-at/pkg/model"
)

var systemProcess []model.ProcessConfig

func processMoveUp(g *gocui.Gui, v *gocui.View) error {
	saveVisualData(g, v)
	gui.CursorUp(g, v)
	getSelectedProcess(g, v)
	refereshGui(g)
	return nil
}

func processMoveDown(g *gocui.Gui, v *gocui.View) error {
	saveVisualData(g, v)
	gui.CursorDown(g, v)
	getSelectedProcess(g, v)
	refereshGui(g)
	return nil
}

func saveVisualData(g *gocui.Gui, v *gocui.View) error {
    if(config.SelectedProcess!=nil) {
        config.SelectedProcess.Text="AAAAAAAAAAAA"
    }
	// ev, err := g.View(model.EDITOR_VIEW)
	// if err == nil {
	// 	vb := ev.ViewBuffer()
        // config.SelectedProcess.Text = vb
	// 	return nil
	// }
	return nil
}
func refereshGui(g *gocui.Gui) error {
	gui.DrawSideGUi(g, config, false)
	gui.DrawMainGui(g, config, false)
	gui.DrawProcessGui(g, config, false)
	gui.DrawEditorGui(g, config, false)
	return nil
}
func addProcessAction(g *gocui.Gui, v *gocui.View) error {

	config.ListProcess = append(config.ListProcess)
	text := gui.GetSelectedText(g, v)
	var selected *model.ProcessConfig
	for _, item := range systemProcess {
		re := regexp.MustCompile(`^\d*\ \-\ `)
		text = re.ReplaceAllString(text, "")
		if item.Title == text {
			selected = &item
			break
		}
	}
	gui.CloseMenuDialog(g, v)
	if selected != nil {
		selected.Name = selected.Title
		item, index := addProcess(selected)
		gui.SetCursorLine(g, model.SIDE_VIEW, index)
		config.SelectedProcess = item
		refereshGui(g)
	}
	return nil
}

func showMenuAddProcess(g *gocui.Gui, v *gocui.View) error {
	systemProcess, _ = command.GetListProcess()
	litem := make([]string, 0)
	for _, l := range systemProcess {
		litem = append(litem, l.Title)
	}
	gui.ShowMenuDiaLog(g, litem, addProcessAction)
	return nil
}

func getSelectedProcess(g *gocui.Gui, v *gocui.View) error {
	text := gui.GetSelectedText(g, v)
	for _, item := range config.ListProcess {
		if item.Name == text {
			config.SelectedProcess = &item
			break
		}
	}
	return nil
}

func deleteSeletedItem(g *gocui.Gui, v *gocui.View) error {
	getSelectedProcess(g, v)
	if config.SelectedProcess != nil {
		removeProcess(config.SelectedProcess)
		refereshGui(g)
	}
	return nil
}

func showDelProcess(g *gocui.Gui, v *gocui.View) error {
	text := gui.GetSelectedText(g, v)
	if text != "" {
		text = fmt.Sprintf("Delete %s?", text)
		gui.ShowDialog(g, v, text, deleteSeletedItem)
	}
	return nil
}

func deleteEditor(g *gocui.Gui, v *gocui.View) error {
	log.Println("aasdas")
	if config.SelectedProcess != nil {
		config.SelectedProcess.Text = ""
		gui.DrawEditorGui(g, config, false)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	command.SaveJSON(config)
	return gui.Quit(g, v)
}
