// Package shared provides ...
package model

import (
	"github.com/jroimartin/gocui"
)

// Package main provides ...
type Point struct {
	x, y int
}
type ProcessConfig struct {
	Pid  int
	Name string
	// current title of pid
	Title  string
	Time   int
	Text   string
	Points []Point
}

type AppConfig struct {
	Status          string
	Message         string
	ListProcess     []ProcessConfig
	SelectedProcess *ProcessConfig
}

const SIDE_VIEW = "side"
const MAIN_VIEW = "main"
const PROCESS_VIEW = "process"
const EDITOR_VIEW = "editor"

const MSG_VIEW = "msg"
const MENU_VIEW = "menu"

const DATA_PATH = "data.json"

type DialogHandler func(g *gocui.Gui, v *gocui.View) error

type ButtonWidget struct {
	name    string
	x, y    int
	w       int
	label   string
	handler func(g *gocui.Gui, v *gocui.View) error
}

func NewButtonWidget(name string, x, y int, label string, handler func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
	return &ButtonWidget{name: name, x: x, y: y, w: len(label) + 1, label: label, handler: handler}
}

// Binding - a keybinding mapping a key and modifier to a handler. The keypress
// is only handled if the given view has focus, or handled globally if the view
// is ""
type Binding struct {
	ViewName    string
	Contexts    []string
	Handler     DialogHandler
	Key         interface{} // FIXME: find out how to get `gocui.Key | rune`
	Modifier    gocui.Modifier
	Description string
	Alternative string
}
