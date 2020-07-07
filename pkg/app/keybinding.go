// Package main provides ...
package app

import (
	// "github.com/nsf/termbox-go"
	"fmt"
	"log"
	"strconv"

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
			Handler:  nextView,
		},
		{
			ViewName:    model.SIDE_VIEW,
			Key:         gocui.KeyTab,
			Handler:     nextView,
			Description: "Focus Next",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'j',
			Handler:     processMoveDown,
			Description: "MoveDown",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'k',
			Handler:     processMoveUp,
			Description: "MoveUp",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'a',
			Handler:     showMenuAddProcess,
			Description: "Add Process",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'x',
			Handler:     showDelProcess,
			Description: "Delete Process",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'c',
			Handler:     showChangeWindowID,
			Description: "ChangeProcess",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         't',
			Handler:     showChangeTime,
			Description: "Change Process Time",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'f',
			Handler:     focusWindow,
			Description: "Focus on process Window",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         's',
			Handler:     showStartRun,
			Description: "Start Task",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'e',
			Handler:     showStopAction,
			Description: "End Task",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'w',
			Handler:     saveDataAction,
			Description: "Write Data",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Handler:     nil,
			Description: "1-9 Move to position",
		},
		{
			ViewName:    model.SIDE_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'H',
			Handler:     showHelp,
			Description: "Show Help",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         gocui.KeyTab,
			Handler:     nextView,
			Description: "Focus Next",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'h',
			Handler:     startHookPoints,
			Description: "Start hook points",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         's',
			Handler:     showStartRun,
			Description: "Start Task",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'e',
			Handler:     showStopAction,
			Description: "End Task",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'r',
			Handler:     showResetProcessData,
			Description: "Reset",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         't',
			Handler:     showChangeTime,
			Description: "Change Process Time",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'w',
			Handler:     saveDataAction,
			Description: "Write Data",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'k',
			Handler:     gui.CursorUp,
			Description: "MoveUp",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'j',
			Handler:     gui.CursorDown,
			Description: "MoveDown",
		},
		{
			ViewName:    model.PROCESS_VIEW,
			Modifier:    gocui.ModNone,
			Key:         'H',
			Handler:     showHelp,
			Description: "Show Help",
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
			Handler:  gui.OkDialog,
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
			ViewName:    model.EDITOR_VIEW,
			Modifier:    gocui.ModNone,
			Key:         gocui.KeyCtrlSpace,
			Handler:     deleteEditor,
			Description: "Editor Clear Editor",
		},
		{
			ViewName:    model.EDITOR_VIEW,
			Modifier:    gocui.ModNone,
			Key:         gocui.KeyCtrlB,
			Handler:     nextView,
			Description: "Editor - Focus process list",
		},
		{
			ViewName:    model.EDITOR_VIEW,
			Modifier:    gocui.ModNone,
			Key:         gocui.KeyCtrlA,
			Handler:     clipboardData,
			Description: "Editor - Get data from clipboard",
		},
		{
			ViewName:    model.EDITOR_VIEW,
			Modifier:    gocui.ModNone,
			Key:         ' ',
			Handler:     nil,
			Description: "ctrl+shift +q stop task global",
		},
		{
			ViewName: "",
			Modifier: gocui.ModNone,
			Key:      gocui.KeyCtrlC,
			Handler:  quit,
		},
	}
	for _, b := range bindings {
		if b.Handler != nil {
			if err := g.SetKeybinding(b.ViewName, b.Key, b.Modifier, b.Handler); err != nil {
				log.Panicf("Error bindings: %s %s", b.ViewName, b.Key)
				return err
			}
		}

	}

	for i := 1; i <= 9; i++ {
		key := []rune(strconv.Itoa(i))[0]
		if err := g.SetKeybinding(model.MENU_VIEW, key, gocui.ModNone, gui.MoveByKey(i-1)); err != nil {
			return err
		}
	}
	for i := 1; i <= 9; i++ {
		key := []rune(strconv.Itoa(i))[0]
		if err := g.SetKeybinding(model.SIDE_VIEW, key, gocui.ModNone, gui.MoveByKey(i-1)); err != nil {
			return err
		}
	}
	return nil
}

var keyMapReversed = map[gocui.Key]string{
	gocui.KeyF1:         "f1",
	gocui.KeyF2:         "f2",
	gocui.KeyF3:         "f3",
	gocui.KeyF4:         "f4",
	gocui.KeyF5:         "f5",
	gocui.KeyF6:         "f6",
	gocui.KeyF7:         "f7",
	gocui.KeyF8:         "f8",
	gocui.KeyF9:         "f9",
	gocui.KeyF10:        "f10",
	gocui.KeyF11:        "f11",
	gocui.KeyF12:        "f12",
	gocui.KeyInsert:     "insert",
	gocui.KeyDelete:     "delete",
	gocui.KeyHome:       "home",
	gocui.KeyEnd:        "end",
	gocui.KeyPgup:       "pgup",
	gocui.KeyPgdn:       "pgdown",
	gocui.KeyArrowUp:    "▲",
	gocui.KeyArrowDown:  "▼",
	gocui.KeyArrowLeft:  "◄",
	gocui.KeyArrowRight: "►",
	gocui.KeyTab:        "tab",        // ctrl+i
	gocui.KeyEnter:      "enter",      // ctrl+m
	gocui.KeyEsc:        "esc",        // ctrl+[, ctrl+3
	gocui.KeyBackspace:  "backspace",  // ctrl+h
	gocui.KeyCtrlSpace:  "ctrl+space", // ctrl+~, ctrl+2
	gocui.KeyCtrlSlash:  "ctrl+/",     // ctrl+_
	gocui.KeySpace:      "space",
	gocui.KeyCtrlA:      "ctrl+a",
	gocui.KeyCtrlB:      "ctrl+b",
	gocui.KeyCtrlC:      "ctrl+c",
	gocui.KeyCtrlD:      "ctrl+d",
	gocui.KeyCtrlE:      "ctrl+e",
	gocui.KeyCtrlF:      "ctrl+f",
	gocui.KeyCtrlG:      "ctrl+g",
	gocui.KeyCtrlJ:      "ctrl+j",
	gocui.KeyCtrlK:      "ctrl+k",
	gocui.KeyCtrlL:      "ctrl+l",
	gocui.KeyCtrlN:      "ctrl+n",
	gocui.KeyCtrlO:      "ctrl+o",
	gocui.KeyCtrlP:      "ctrl+p",
	gocui.KeyCtrlQ:      "ctrl+q",
	gocui.KeyCtrlR:      "ctrl+r",
	gocui.KeyCtrlS:      "ctrl+s",
	gocui.KeyCtrlT:      "ctrl+t",
	gocui.KeyCtrlU:      "ctrl+u",
	gocui.KeyCtrlV:      "ctrl+v",
	gocui.KeyCtrlW:      "ctrl+w",
	gocui.KeyCtrlX:      "ctrl+x",
	gocui.KeyCtrlY:      "ctrl+y",
	gocui.KeyCtrlZ:      "ctrl+z",
	gocui.KeyCtrl4:      "ctrl+4", // ctrl+\
	gocui.KeyCtrl5:      "ctrl+5", // ctrl+]
	gocui.KeyCtrl6:      "ctrl+6",
	gocui.KeyCtrl8:      "ctrl+8",
}

var keymap = map[string]interface{}{
	"<c-a>":       gocui.KeyCtrlA,
	"<c-b>":       gocui.KeyCtrlB,
	"<c-c>":       gocui.KeyCtrlC,
	"<c-d>":       gocui.KeyCtrlD,
	"<c-e>":       gocui.KeyCtrlE,
	"<c-f>":       gocui.KeyCtrlF,
	"<c-g>":       gocui.KeyCtrlG,
	"<c-h>":       gocui.KeyCtrlH,
	"<c-i>":       gocui.KeyCtrlI,
	"<c-j>":       gocui.KeyCtrlJ,
	"<c-k>":       gocui.KeyCtrlK,
	"<c-l>":       gocui.KeyCtrlL,
	"<c-m>":       gocui.KeyCtrlM,
	"<c-n>":       gocui.KeyCtrlN,
	"<c-o>":       gocui.KeyCtrlO,
	"<c-p>":       gocui.KeyCtrlP,
	"<c-q>":       gocui.KeyCtrlQ,
	"<c-r>":       gocui.KeyCtrlR,
	"<c-s>":       gocui.KeyCtrlS,
	"<c-t>":       gocui.KeyCtrlT,
	"<c-u>":       gocui.KeyCtrlU,
	"<c-v>":       gocui.KeyCtrlV,
	"<c-w>":       gocui.KeyCtrlW,
	"<c-x>":       gocui.KeyCtrlX,
	"<c-y>":       gocui.KeyCtrlY,
	"<c-z>":       gocui.KeyCtrlZ,
	"<c-~>":       gocui.KeyCtrlTilde,
	"<c-2>":       gocui.KeyCtrl2,
	"<c-3>":       gocui.KeyCtrl3,
	"<c-4>":       gocui.KeyCtrl4,
	"<c-5>":       gocui.KeyCtrl5,
	"<c-6>":       gocui.KeyCtrl6,
	"<c-7>":       gocui.KeyCtrl7,
	"<c-8>":       gocui.KeyCtrl8,
	"<c-space>":   gocui.KeyCtrlSpace,
	"<c-\\>":      gocui.KeyCtrlBackslash,
	"<c-[>":       gocui.KeyCtrlLsqBracket,
	"<c-]>":       gocui.KeyCtrlRsqBracket,
	"<c-/>":       gocui.KeyCtrlSlash,
	"<c-_>":       gocui.KeyCtrlUnderscore,
	"<backspace>": gocui.KeyBackspace,
	"<tab>":       gocui.KeyTab,
	"<enter>":     gocui.KeyEnter,
	"<esc>":       gocui.KeyEsc,
	"<space>":     gocui.KeySpace,
	"<f1>":        gocui.KeyF1,
	"<f2>":        gocui.KeyF2,
	"<f3>":        gocui.KeyF3,
	"<f4>":        gocui.KeyF4,
	"<f5>":        gocui.KeyF5,
	"<f6>":        gocui.KeyF6,
	"<f7>":        gocui.KeyF7,
	"<f8>":        gocui.KeyF8,
	"<f9>":        gocui.KeyF9,
	"<f10>":       gocui.KeyF10,
	"<f11>":       gocui.KeyF11,
	"<f12>":       gocui.KeyF12,
	"<insert>":    gocui.KeyInsert,
	"<delete>":    gocui.KeyDelete,
	"<home>":      gocui.KeyHome,
	"<end>":       gocui.KeyEnd,
	"<pgup>":      gocui.KeyPgup,
	"<pgdown>":    gocui.KeyPgdn,
	"<up>":        gocui.KeyArrowUp,
	"<down>":      gocui.KeyArrowDown,
	"<left>":      gocui.KeyArrowLeft,
	"<right>":     gocui.KeyArrowRight,
}

func GetKeyDisplay(key interface{}) string {
	keyInt := 0

	switch key := key.(type) {
	case rune:
		keyInt = int(key)
	case gocui.Key:
		value, ok := keyMapReversed[key]
		if ok {
			return value
		}
		keyInt = int(key)
	}

	return string(keyInt)
}

func showHelp(g *gocui.Gui, v *gocui.View) error {
	viewName := v.Name()
	listHelpKey := make([]string, 0)
	for _, b := range bindings {
		if b.ViewName != "" && (b.ViewName == viewName || b.ViewName == model.EDITOR_VIEW) {
			listHelpKey = append(listHelpKey, fmt.Sprintf("%s %s", GetKeyDisplay(b.Key), b.Description))
		}
	}
	gui.ShowMenuDiaLog(g, v, listHelpKey, gui.CloseDialog)
	return nil

}
