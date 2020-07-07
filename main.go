package main

import (
	"fmt"
	"os/signal"
	"syscall"

	// "io/ioutil"
	"log"
	"os"

	// "github.com/nsf/termbox-go"
	"github.com/go-vgo/robotgo"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/app"
	"github.com/windwp/go-at/pkg/command"
	"github.com/windwp/go-at/pkg/gui"
	"github.com/windwp/go-at/pkg/model"
)

var (
	config *model.AppConfig
)

func layout(g *gocui.Gui) error {
	gui.SetUpGui(g, config)
	return nil
}
func onKill() {
	c := make(chan os.Signal)
	log.Println("OnKill")
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		<-c
		log.Println("Stop task")
		go command.EndTask()
	}()
}
func saveAppData() {
	command.SaveJSON(config)
}

func main() {
	f, err := os.OpenFile("./at.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("file log not exist")
	}
	defer f.Close()
	defer saveAppData()
	log.SetOutput(f)
	g, err := gocui.NewGui(gocui.OutputNormal)
	config = app.Setup()
	robotgo.GetMousePos()
	robotgo.TypeStr("")
	onKill()
	command.CLICK_TIME_DURATION = 8
	command.SLEEP_TASK_MILISECOND = 300
	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.InputEsc = true
	if err != nil {
		g.Close()
		log.Panicln(err)
		return
	}
	defer g.Close()
	g.SetManagerFunc(layout)
	if err := app.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
