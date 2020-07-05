
package gui

import (
	"fmt"
	"log"
	"regexp"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

var menuHandler model.DialogHandler
var dialogHandler model.DialogHandler

func ShowMenuDiaLog(g *gocui.Gui, listItem []string, handler model.DialogHandler) error {
	w := 50
	h := 20
	maxX, maxY := g.Size()
	if w > maxX {
		w = maxX - 10
	}
	if v, err := g.SetView(model.MENU_VIEW, maxX/2-w/2, maxY/2-h/2, maxX/2+w/2, maxY/2+h/2); err != nil {
		v.Title = "Dialog"
		v.Highlight = true
		if err != gocui.ErrUnknownView {
			return err
		}
		for i, l := range listItem {
			fmt.Fprintf(v, menu_item_format, i+1, l)
		}
		if _, err := g.SetCurrentView(model.MENU_VIEW); err != nil {
			log.Panic("can't set view")
			return err
		}
	}
	return nil
}

func GetSelectedText(g *gocui.Gui, v *gocui.View) string {
	var text string
	var err error
	_, cy := v.Cursor()
	if text, err = v.Line(cy); err != nil {
		return ""
	}
	re := regexp.MustCompile(`^\d*\ `)
	text = re.ReplaceAllString(text, "")
	return text

}

func CloseMenuDialog(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(model.MENU_VIEW); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
		return err
	}
	menuHandler = nil
	return nil
}

func OkDialog(g *gocui.Gui, v *gocui.View) error {
	if dialogHandler != nil {
		dialogHandler(g, v)
	}
	CloseDialog(g, v)
	return nil
}

func CloseDialog(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(model.MSG_VIEW); err != nil {
		return err
	}
	if _, err := g.SetCurrentView(model.SIDE_VIEW); err != nil {
		return err
	}
	dialogHandler = nil
	return nil
}

func ShowDialog(
	g *gocui.Gui,
	v *gocui.View,
	message string, handler model.DialogHandler,
) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(model.MSG_VIEW, maxX/2-20, maxY/2-5, maxX/2+20, maxY/2); err != nil {
		v.Title = "Dialog"
		if err != gocui.ErrUnknownView {
			log.Panic("stupi")
			return err
		}
		fmt.Fprintln(v, message)
		fmt.Fprintln(v, " ")
		text := "[Y]es  |   [N]o"
		cursorX := 25 - len(text)
		for i := 0; i < cursorX; i++ {
			text = " " + text
		}

		v.SetCursor(cursorX+1, 2)
		fmt.Fprintf(v, "%s\n", text)
		if _, err := g.SetCurrentView(model.MSG_VIEW); err != nil {
			log.Panic("can't set view")
			return err
		}
		dialogHandler = handler
	}

	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func CursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		n, _ := v.Line(cy + 1)
		if n != "" {
			if err := v.SetCursor(cx, cy+1); err != nil {
				ox, oy := v.Origin()
				if err := v.SetOrigin(ox, oy+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func SetCursorLine(g *gocui.Gui, viewName string,index int) error {
	v,err:=g.View(viewName)
	if v != nil&& err==nil {
		ox, oy := v.Origin()
		cx, _ := v.Cursor()
		if err := v.SetCursor(cx, index); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// ColoredString takes a string and a colour attribute and returns a colored
// string with that attribute
func ColoredString(str string, colorAttributes ...color.Attribute) string {
	colour := color.New(colorAttributes...)
	return ColoredStringDirect(str, colour)
}

// ColoredStringDirect used for aggregating a few color attributes rather than
// just sending a single one
func ColoredStringDirect(str string, colour *color.Color) string {
	return colour.SprintFunc()(fmt.Sprint(str))
}