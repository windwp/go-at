// Package main provides ...
package app
import (
	// "github.com/nsf/termbox-go"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/windwp/go-at/pkg/model"
)

// Package main provides ...
func Keybindings(g *gocui.Gui) error {

    bindings := []*model.Binding{
        {
            ViewName  :model.MAIN_VIEW,
            Key: gocui.KeyCtrlSpace,
            Handler:nextView,
        },
        {
            ViewName  :model.SIDE_VIEW,
            Key: gocui.KeyCtrlSpace,
            Handler:nextView,
        },

        {
            ViewName  :model.SIDE_VIEW,
            Key:'j',
            Handler:cursorDown,
        },
        {
            ViewName  :model.SIDE_VIEW,
            Key:'k',
            Handler:cursorUp,
        },
        {
            ViewName  :model.SIDE_VIEW,
            Key:'a',
            Handler:addProcess,
        },
        {
            ViewName  :model.SIDE_VIEW,
            Key:'x',
            Handler:delItem,
        },


        {
            ViewName  :model.MSG_VIEW,
            Key:'y',
            Handler:okDialog,
        },
        {
            ViewName  :model.MSG_VIEW,
            Key:'n',
            Handler:closeDialog,
        },
        {
            ViewName  :model.MSG_VIEW,
            Key:gocui.KeyEnter,
            Handler:okDialog,
        },


        {
            ViewName  :model.PROCESS_VIEW,
            Key:'j',
            Handler:cursorDown,
        },
        {
            ViewName  :model.PROCESS_VIEW,
            Key:'k',
            Handler:cursorUp,
        },
    }

    for _, b := range bindings {
        if err := g.SetKeybinding(b.ViewName,b.Key,gocui.ModNone,b.Handler); err != nil {
            log.Panicf("Error bindings: %s %s",b.ViewName,b.Key)
            return err
        }
    
    }
	return nil
}

