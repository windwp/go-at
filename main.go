package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"os"

	// "github.com/nsf/termbox-go"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/app"
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

func main() {
	os.Remove("./at.log")
	f, err := os.OpenFile("./at.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file log not exist")
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Start")

	g, err := gocui.NewGui(gocui.OutputNormal)

	config = app.Setup()
	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

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
