// Package command provides ...
package command

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/windwp/go-at/pkg/model"
)

func CheckProcess(config *model.AppConfig) error {
	psList, _ := GetListProcess()
	for _, p := range config.ListProcess {
		isExist := false
		for _, cP := range psList {
			if p.Pid == cP.Pid {
				isExist = true
			}
		}
		if isExist == false {
			return errors.New(fmt.Sprintf("Not found %s", p.Title))
		}

	}
	return nil
}
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func doTask(p *model.ProcessConfig, pIndex, pointIndex, textIndex int, time int64) (pI, tI int) {
	if len(p.Text) <= textIndex {
		textIndex = 0
	}
	var key = p.Text[textIndex:1]
	robotgo.TypeStr(key)
	return pointIndex, textIndex
}

var endEvent = make(chan int)
var startChan = make(chan int)

func EndTask() {
	endEvent <- 1
}
func StartTask(config *model.AppConfig) error {
	err := CheckProcess(config)
	if err != nil {
		return err
	}
	pIndex := 0
	pointIndex := 0
	textIndex := 0
	startTime := makeTimestamp()
	var currentProcess *model.ProcessConfig
	go func() {
		if currentProcess == nil {
			if len(config.ListProcess) <= pIndex {
				pIndex = 0
			}
			currentProcess := &config.ListProcess[pIndex]
			FocusWindow(currentProcess)
			time.Sleep(1 * time.Second)

		}
		pointIndex, textIndex = doTask(currentProcess, pIndex, pointIndex, textIndex, startTime)
		time.Sleep(1 * time.Second)
		cTime := makeTimestamp()
		if startTime+(int64)(currentProcess.Time) > cTime {
			pIndex += 1
			currentProcess = nil
		}
		for {
			select {
			case endEvent <- 1:
				return
			}
		}
	}()

	return nil
}
