// Package main provides ...
package app

import (
	// "github.com/nsf/termbox-go"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/gui"
	"github.com/windwp/go-at/pkg/model"
)

// Package main provides ...
func Keybindings(g *gocui.Gui) error {

	bindings := []*model.Binding{
		{
			ViewName: model.MAIN_VIEW,
			Key:      gocui.KeyTab,
			Handler:  gui.NextView,
		},
		{
			ViewName: model.SIDE_VIEW,
			Key:      gocui.KeyTab,
			Handler:  gui.NextView,
		},

		{
			ViewName: model.PROCESS_VIEW,
			Key:      gocui.KeyTab,
			Handler:  gui.NextView,
		},
		{
			ViewName: model.EDITOR_VIEW,
			Key:      gocui.KeyTab,
			Handler:  gui.NextView,
		},
		{
			ViewName: model.SIDE_VIEW,
			Key:      'j',
			Handler:  processMoveDown,
		},
		{
			ViewName: model.SIDE_VIEW,
			Key:      'k',
			Handler:  processMoveUp,
		},
		{
			ViewName: model.SIDE_VIEW,
			Key:      'a',
			Handler:  showMenuAddProcess,
		},
		{
			ViewName: model.SIDE_VIEW,
			Key:      'x',
			Handler:  showDelProcess,
		},

		{
			ViewName: model.MSG_VIEW,
			Key:      'y',
			Handler:  gui.OkDialog,
		},
		{
			ViewName: model.MSG_VIEW,
			Key:      'n',
			Handler:  gui.OkDialog,
		},
		{
			ViewName: model.MSG_VIEW,
			Key:      gocui.KeyEnter,
			Handler:  gui.OkDialog,
		},

		{
			ViewName: model.MENU_VIEW,
			Key:      'j',
			Handler:  gui.CursorDown,
		},
		{
			ViewName: model.MENU_VIEW,
			Key:      'k',
			Handler:  gui.CursorUp,
		},
		{
			ViewName: model.MENU_VIEW,
			Key:      gocui.KeyEnter,
			Handler:  addProcessAction,
		},
		{
			ViewName: model.MENU_VIEW,
			Key:      'q',
			Handler:  gui.CloseMenuDialog,
		},

		{
			ViewName: model.MENU_VIEW,
			Key:      gocui.KeyEsc,
			Handler:  gui.CloseMenuDialog,
		},
        {
            ViewName: model.EDITOR_VIEW,
            Key:      gocui.KeyCtrlSpace,
            Handler:  deleteEditor,
        },
		{
			ViewName: "",
			Key:      gocui.KeyCtrlC,
			Handler:  quit,
		},
	}

	for _, b := range bindings {
		if err := g.SetKeybinding(b.ViewName, b.Key, gocui.ModNone, b.Handler); err != nil {
			log.Panicf("Error bindings: %s %s", b.ViewName, b.Key)
			return err
		}

	}
	return nil
}
