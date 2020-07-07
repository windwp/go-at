package app

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
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

func nextView(g *gocui.Gui, v *gocui.View) error {
	saveVisualData(g, v)
	return gui.NextView(g, v)
}
func saveVisualData(g *gocui.Gui, v *gocui.View) error {
	if config.SelectedProcess != nil {
		ev, err := g.View(model.EDITOR_VIEW)
		if err == nil {
			ev.Rewind()
			vb := ev.ViewBuffer()
			vb = strings.Trim(vb, "\n")
			vb = strings.Trim(vb, " ")
			config.SelectedProcess.Text = vb
			return nil
		}
	}
	return nil
}
func getSysemSelectProcess(g *gocui.Gui, v *gocui.View) *model.ProcessConfig {
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
	return selected
}
func changeProcessAction(g *gocui.Gui, v *gocui.View) error {
	selected := getSysemSelectProcess(g, v)
	gui.CloseMenuDialog(g, v)
	if selected != nil {
		selected.Name = selected.Title
		config.SelectedProcess.Pid = selected.Pid
		config.SelectedProcess.Wid = selected.Wid
		config.SelectedProcess.Name = selected.Name
		config.SelectedProcess.Title = selected.Title
		refereshGui(g)
	}
	return nil
}

func showChangeWindowID(g *gocui.Gui, v *gocui.View) error {
	systemProcess, _ = command.GetListProcess()
	litem := make([]string, 0)
	for _, l := range systemProcess {
		litem = append(litem, l.Title)
	}
	gui.ShowMenuDiaLog(g, v, litem, changeProcessAction)
	return nil
}

func changeTimeAction(g *gocui.Gui, v *gocui.View, text string) error {
	if config.SelectedProcess != nil {
		number, err := strconv.Atoi(text)
		if err == nil {
			config.SelectedProcess.Time = number
			gui.DrawProcessGui(g, config, false)
		}

	}
	return nil
}

func showChangeTime(g *gocui.Gui, v *gocui.View) error {
	gui.ShowInputDialog(g, v, "Change Time", changeTimeAction)
	return nil
}
func refereshGui(g *gocui.Gui) error {
	gui.DrawMainGui(g, config, false)
	gui.DrawSideGUi(g, config, false)
	gui.DrawProcessGui(g, config, false)
	gui.DrawEditorGui(g, config, false)
	gui.DrawStatusGui(g, config, false)
	return nil
}
func addProcessAction(g *gocui.Gui, v *gocui.View) error {
	selected := getSysemSelectProcess(g, v)
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
	gui.ShowMenuDiaLog(g, v, litem, addProcessAction)
	return nil
}

func getSelectedProcess(g *gocui.Gui, v *gocui.View) error {
	text := gui.GetSelectedText(g, v)
	for i, item := range config.ListProcess {
		if item.Name == text {
			config.SelectedProcess = &config.ListProcess[i]
			return nil
		}
	}
	return nil
}

func onEndTask(g *gocui.Gui, v *gocui.View) error{
    <-command.WaitTask()
    stopRunAction(g,v)
    return nil
}
func startRunAction(g *gocui.Gui, v *gocui.View) error {
	if config.Status == model.S_IDLE {
		err := command.StartTask(config, false)
		if err != nil {
			SetMessage(err.Error(), g)
			return err
		}
        go onEndTask(g,v)
		config.Status = model.S_RUNNING
		refereshGui(g)
	}
	return nil
}

func stopRunAction(g *gocui.Gui, v *gocui.View) error {
	if config.Status == model.S_RUNNING {
        command.EndTask()
        config.Status = model.S_IDLE
        SetMessage("stop", g)
	}
	return nil
}
func showStartRun(g *gocui.Gui, v *gocui.View) error {
	if config.Status == model.S_IDLE {
		saveVisualData(g, v)
		gui.ShowDialog(g, v, "Start ?", startRunAction)
	} else {
		gui.ShowDialog(g, v, "Can't start Now", gui.CloseDialog)
	}
	return nil
}

func clipboardData(g *gocui.Gui, v *gocui.View) error {
	clipboard, err := robotgo.ReadAll()
	if err == nil {
		config.SelectedProcess.Text = clipboard
		refereshGui(g)
	} else {

		log.Panic("Clip board error")
	}
	return nil
}
func showStopAction(g *gocui.Gui, v *gocui.View) error {
	if config.Status == model.S_RUNNING {
		gui.ShowDialog(g, v, "Stop? ", stopRunAction)
	} else {
		gui.ShowDialog(g, v, "Can't start Now", gui.CloseDialog)
	}
	return nil
}
func deleteSeletedProcess(g *gocui.Gui, v *gocui.View) error {
	getSelectedProcess(g, v)
	if config.SelectedProcess != nil {
		removeProcess(config.SelectedProcess)
		refereshGui(g)
	} else {
		log.Panic(config.SelectedProcess)

	}
	return nil
}

func showDelProcess(g *gocui.Gui, v *gocui.View) error {
	text := gui.GetSelectedText(g, v)
	if text != "" {
		text = fmt.Sprintf("Delete %s?", text)
		gui.ShowDialog(g, v, text, deleteSeletedProcess)
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

func focusWindow(g *gocui.Gui, v *gocui.View) error {
	if config.SelectedProcess != nil {
		err := command.FocusWindow(config.SelectedProcess)
		if err != nil {
			gui.ShowDialog(g, v, err.Error(), gui.CloseDialog)
		}
	}
	return nil
}
func saveDataAction(g *gocui.Gui, v *gocui.View) error {
	command.SaveJSON(config)
	SetMessage("Data Saved", g)
	gui.DrawStatusGui(g, config, false)
	return nil
}

func showResetProcessData(g *gocui.Gui, v *gocui.View) error {
	tmp := func(g *gocui.Gui, v *gocui.View) error {
		if config.SelectedProcess != nil {
			model.ResetProcess(config.SelectedProcess)
		}
		refereshGui(g)
		return nil
	}
	gui.ShowDialog(g, v, "Reset data", tmp)
	return nil
}
func startHookPoints(g *gocui.Gui, v *gocui.View) error {
	update := make(chan int)
	end := make(chan int)
	err := startHookKeyBoard(update, end)
	if err != nil {
		gui.ShowDialog(g, v, err.Error(), gui.CloseDialog)
		return nil
	}
	go func() {
		for {
			select {
			case <-update:
				refereshGui(g)
			case <-end:
				refereshGui(g)
				return
			}
		}
	}()
	refereshGui(g)
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	command.SaveJSON(config)
	return gui.Quit(g, v)
}
