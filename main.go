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
	// maxX, maxY := g.Size()
	gui.SetUpSideGui(g, config)
    gui.SetUpMainGui(g,config)
	// if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
	// 	if err != gocui.ErrUnknownView {
	// 		return err
	// 	}

	// 	b, err := ioutil.ReadFile("main.go")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Fprintf(v, "%s", b)
	// 	v.Editable = true
	// 	v.Wrap = true
	// 	if _, err := g.SetCurrentView("side"); err != nil {
	// 		return err
	// 	}
	// }

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

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	if err != nil {
		g.Close()
		log.Panicln(err)
		return
	}
	defer g.Close()
	g.Cursor = true
	g.SetManagerFunc(layout)

	if err := app.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
