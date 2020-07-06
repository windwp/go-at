package app

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/jroimartin/gocui"
	hook "github.com/robotn/gohook"
	"github.com/windwp/go-at/pkg/command"
	"github.com/windwp/go-at/pkg/model"
)

// Package main provides ...
var config *model.AppConfig

func Setup() *model.AppConfig {
	pNum := 10
	config = &model.AppConfig{
		Status:      model.S_IDLE,
		Message:     "OK",
		ListProcess: make([]model.ProcessConfig, 0, pNum),
	}
	jsonConfig, err := command.LoadJson()
	if err == nil && jsonConfig != nil {
		config = jsonConfig
	}

	config.Status = model.S_IDLE
	config.Message = ""
	if len(config.ListProcess) > 0 {
		config.SelectedProcess = &config.ListProcess[0]
	}

	return config
}

func addProcess(p *model.ProcessConfig) (*model.ProcessConfig, int) {
	isvalid := true
	for index, i := range config.ListProcess {
		if i.Pid == p.Pid {
			isvalid = false
			return &i, index
		}
	}
	if isvalid {
		config.ListProcess = append(config.ListProcess, *p)
		log.Printf(" length %d", len(config.ListProcess))
	}
	return p, len(config.ListProcess) - 1
}

func removeProcess(p *model.ProcessConfig) int {
	if p != nil {
		nProcess := make([]model.ProcessConfig, 0)
		index := 0
		for i := 0; i < len(config.ListProcess); i++ {
			item := config.ListProcess[i]
			if item.Pid != p.Pid {
				nProcess = append(nProcess, item)
			} else {
				index = i
			}
		}
		config.ListProcess = nProcess
		if index < len(config.ListProcess) {
			config.SelectedProcess = &config.ListProcess[index]
		}
	}
	return 0
}

func SetMessage(message string,g *gocui.Gui) {
	config.Message = message
	go func() {
		time.Sleep(3 * time.Second)
		config.Message = ""
        refereshGui(g)
	}()
    
}
func addPoint(p *model.ProcessConfig, x, y int) {
	if p == nil {
		return
	}
	newPoint := &model.Point{X: x, Y: y}
	p.Points = append(p.Points, *newPoint)
}

func updateProcess(p *model.ProcessConfig) {
	for _, i := range config.ListProcess {
		if i.Pid == p.Pid {
			(&i).Text = p.Text
			(&i).Time = p.Time
			(&i).Points = p.Points
			break
		}
	}
}

func startHookKeyBoard(update, end chan int) error {
	if config.Status == model.S_IDLE {
		go func(update, end chan int) {
			robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
				config.Status = model.S_IDLE
				time.Sleep(time.Second * 1)
				update <- 0
				end <- 0
				robotgo.EventEnd()
			})
			robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
				x, y := robotgo.GetMousePos()
				a := strconv.Itoa(x)
				if len(a) < 5 {
					addPoint(config.SelectedProcess, x, y)
					update <- 0
				}
			})
			s := robotgo.EventStart()
			<-robotgo.EventProcess(s)
		}(update, end)
		config.Status = model.S_HOOK
		return nil
	} else {
		return errors.New("Server is running")
	}
}
