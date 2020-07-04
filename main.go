
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
    // "github.com/nsf/termbox-go"
	"github.com/jroimartin/gocui"
    "github.com/windwp/go-at/pkg/app"
)

var (
    config app.AppConfig
)


func getLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2-10, maxX/2+30, maxY/2+30); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l+"lxla")
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}





func saveMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func saveVisualMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func updateView(g *gocui.Gui, v *gocui.View) error {
        v, err := g.View("side")
        if err != nil {
            return err
        }
        v.Clear()
        fmt.Fprintln(v, "lala")
        return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
        for _,c := range config.ListProcess {
            fmt.Fprintln(v,c.Name) 
        }
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")

	}

    
	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		b, err := ioutil.ReadFile("main.go")
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(v, "%s", b)
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("side"); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
    config=app.Setup()

	if err != nil {
        g.Close()
        log.Println("error my haha hihi")
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
