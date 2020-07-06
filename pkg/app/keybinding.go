// Package main provides ...
package app

import (
	// "github.com/nsf/termbox-go"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/gui"
	"github.com/windwp/go-at/pkg/model"
)

var bindings []*model.Binding

// Package main provides ...

func Keybindings(g *gocui.Gui) error {
	bindings = []*model.Binding{
		{
			ViewName: model.MAIN_VIEW,
			Key:      gocui.KeyTab,
			Modifier: gocui.ModNone,
			Handler:  gui.NextView,
		},
		{
			ViewName:    model.SIDE_VIEW,
			Key:         gocui.KeyTab,
			Handler:     gui.NextView,
			Description: "Tab - Focus Next",
		},

		{
			ViewName: model.EDITOR_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyTab,
			Handler:  gui.NextView,
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'j',
			Handler:     processMoveDown,
			Description: "j - MoveDown",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'k',
			Handler:     processMoveUp,
			Description: "k - MoveUp",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'a',
			Handler:     showMenuAddProcess,
			Description: "a - Add Process",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'x',
			Handler:     showDelProcess,
			Description: "x - Delete Process",
		},

		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         't',
			Handler:     showChangeTime,
			Description: "t - Change Process Time",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'e',
			Handler:     gui.FocusView(model.EDITOR_VIEW),
			Description: "e - Focus editor view",
		},

		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'f',
			Handler:     focusWindow,
			Description: "f - focus on process",
		},

		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         's',
			Handler:     showStartRun,
			Description: "s - Write Data",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'w',
			Handler:     saveDataAction,
			Description: "w - Write Data",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'H',
			Handler:     showHelp,
			Description: "H - Show Help",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         gocui.KeyTab,
			Handler:     gui.NextView,
			Description: "Tab - Focus Next",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'p',
			Handler:     startHookPoints,
			Description: "f start hook points",
		},

		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'H',
			Handler:     showHelp,
			Description: "H - Show Help",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'k',
			Handler:     gui.CursorUp,
			Description: "k - MoveUp",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'j',
			Handler:     gui.CursorDown,
			Description: "j - MoveDown",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         't',
			Handler:     showChangeTime,
			Description: "t - Change Process Time",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'r',
			Handler:     showResetProcessData,
			Description: "r - Reset",
		},
        {
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         's',
			Handler:     showStartRun,
			Description: "s - Start Task",
		},

		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'w',
			Handler:     saveDataAction,
			Description: "w - Write Data",
		},
		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      'y',
			Handler:  gui.OkDialog,
		},

		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      'n',
			Handler:  gui.CloseDialog,
		},
		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      'q',
			Handler:  gui.CloseDialog,
		},
		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEsc,
			Handler:  gui.CloseDialog,
		},
		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEnter,
			Handler:  gui.OkDialog,
		},

		{
			ViewName: model.MSG_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEsc,
			Handler:  gui.OkInputDialog,
		},
		{
			ViewName: model.INPUT_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEsc,
			Handler:  gui.CloseInputDialog,
		},
		{
			ViewName: model.INPUT_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEnter,
			Handler:  gui.OkInputDialog,
		},
		{
			Modifier: gocui.ModNone,
			ViewName: model.MENU_VIEW,
			Key:      'j',
			Handler:  gui.CursorDown,
		},
		{
			ViewName: model.MENU_VIEW,
			Modifier: gocui.ModNone,
			Key:      'k',
			Handler:  gui.CursorUp,
		},
		{
			ViewName: model.MENU_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEnter,
			Handler:  addProcessAction,
		},
		{
			ViewName: model.MENU_VIEW,
			Modifier: gocui.ModNone,
			Key:      'q',
			Handler:  gui.CloseMenuDialog,
		},
		{
			ViewName: model.MENU_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyEsc,
			Handler:  gui.CloseMenuDialog,
		},
		{
			ViewName: model.EDITOR_VIEW,
			Modifier: gocui.ModNone,
			Key:      gocui.KeyCtrlSpace,
			Handler:  deleteEditor,
		},

		{
			ViewName: "",
			Modifier: gocui.ModNone,
			Key:      gocui.KeyCtrlC,
			Handler:  quit,
		},
	}
	for _, b := range bindings {
		if err := g.SetKeybinding(b.ViewName, b.Key, b.Modifier, b.Handler); err != nil {
			log.Panicf("Error bindings: %s %s", b.ViewName, b.Key)
			return err
		}

	}
	return nil
}

func showHelp(g *gocui.Gui, v *gocui.View) error {
	viewName := v.Name()
	listHelpKey := make([]string, 0)
	for _, b := range bindings {
		if b.ViewName == viewName {
			listHelpKey = append(listHelpKey, fmt.Sprintf("%s", b.Description))
		}
	}
	gui.ShowMenuDiaLog(g,v, listHelpKey, gui.CloseDialog)
	return nil

}
