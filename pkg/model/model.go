// Package shared provides ...
package model

import (
	"github.com/jroimartin/gocui"
)

// Package main provides ...
type Point struct {
	X, Y int
}
type ProcessConfig struct {
	Wid  string
	Pid  int
	Name string
	// current title of pid
	Title      string
	Time       int
	Text       string
	Points     []Point
	PointIndex int
	TextIndex  int
}

type AppStatus string

const (
	S_IDLE    AppStatus = "IDLE"
	S_HOOK    AppStatus = "HOOK"
	S_RUNNING AppStatus = "RUNNING"
	S_PAUSE   AppStatus = "PAUSE"
)

var statusDict = map[AppStatus]string{
	S_IDLE:    "IDLE",
	S_HOOK:    "HOOK",
	S_RUNNING: "RUNNING",
	S_PAUSE:   "PAUSE",
}

func (s AppStatus) String() string {
	a, exist := statusDict[s]
	if exist {
		return a
	} else {
		return "Not Valid"
	}
}

type AppConfig struct {
	Status          AppStatus
	Message         string
	ListProcess     []ProcessConfig
	SelectedProcess *ProcessConfig
}

const SIDE_VIEW = "side"
const MAIN_VIEW = "main"
const STATUS_VIEW = "status"
const PROCESS_VIEW = "process"
const EDITOR_VIEW = "editor"

const MSG_VIEW = "msg"
const INPUT_VIEW = "input"
const MENU_VIEW = "menu"

const DATA_PATH = "data.json"

type DialogHandler func(g *gocui.Gui, v *gocui.View) error
type InputDialogHandler func(g *gocui.Gui, v *gocui.View, value string) error

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

func ResetProcess(p *ProcessConfig) {
	p.Points = make([]Point, 0)
	p.Time = 120
}
