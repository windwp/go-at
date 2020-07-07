// Package command provides ...
package command

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
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
	return time.Now().Unix()
}

// time change for click
var CLICK_TIME_DURATION = 5

// time sleep milisecond on task thread
var SLEEP_TASK_MILISECOND = 500
var KEYPRESS_AFTER_MOUSE_CLICK = []string{"num-", "alt"}
var KEYPRESS_STOP_GLOBAL_HOOK = []string{"q", "ctrl", "shift"}
var endEvent = make(chan int)
var startChan = make(chan int)
var _isRun = false
var _debug = false

func EndTask() {
	if _isRun {
		logText("END")
		endEvent <- 1
		_isRun = false
        hook.End()
	}
}
func WaitTask() chan int {
	return endEvent
}
func logText(text string) {
	if _debug {
		fmt.Println(text)
	} else {
		log.Println(text)
	}
}

func doTask(p *model.ProcessConfig, pIndex, pointIndex, textIndex int, isClick bool) (pI, tI int) {
	if p == nil {
		logText("Process is null")
		return 0, 0
	}
	if isClick == false {
		if len(p.Text) == 0 {
			return pointIndex, textIndex
		}
		if len(p.Text) <= textIndex {
			textIndex = 0
		}
		var key = robotgo.CharCodeAt(p.Text, textIndex)
		textIndex += 1
		logText(fmt.Sprintf("Press Key %d", key))
		robotgo.UnicodeType(uint32(key))
	} else {
		if len(p.Points) == 0 {
			return pointIndex, textIndex
		}
		if len(p.Points) <= pointIndex {
			pointIndex = 0
		}
		point := p.Points[pointIndex]
		pointIndex += 1
		logText("Mouse Click")
		robotgo.MoveMouse(point.X, point.Y)
		robotgo.MoveClick(point.X, point.Y)
		time.Sleep(time.Duration(SLEEP_TASK_MILISECOND) * time.Millisecond)
		// after mouse click press this key
		robotgo.KeyTap(
			KEYPRESS_AFTER_MOUSE_CLICK[0],
			KEYPRESS_AFTER_MOUSE_CLICK[1],
		)
	}
	return pointIndex, textIndex
}
func hookKeyboard() {
	hook.Register(hook.KeyDown, KEYPRESS_STOP_GLOBAL_HOOK, func(e hook.Event) {
		EndTask()
	})

	s := hook.Start()
	<-hook.Process(s)
}
func StartTask(config *model.AppConfig, debug bool) error {
	_debug = debug
	err := CheckProcess(config)
	if err != nil {
		go EndTask()
		return err
	}
	if _isRun {
		go EndTask()
		return errors.New("Task is Running")
	}
	if len(config.ListProcess) <= 0 || config.ListProcess == nil {
		go EndTask()
		return errors.New("No task")
	}
	pIndex := 0
	pointIndex := 0
	textIndex := 0
	startTime := makeTimestamp()
	clickTimeDuration := CLICK_TIME_DURATION
	sleepTimeDuration := SLEEP_TASK_MILISECOND
	isClick := false
	keyTime := makeTimestamp()
	var currentProcess *model.ProcessConfig
	logText(fmt.Sprintf("==== Start ==== %d\n", startTime))
	if debug {
		robotgo.TypeStr("")
	}
	for i := 0; i < len(config.ListProcess); i++ {
		item := &config.ListProcess[i]
		item.PointIndex = 0
		item.TextIndex = 0
	}
	go func() {
		logText("==== Start ====")
		_isRun = true
		go hookKeyboard()
		for {
			select {
			case <-endEvent:
				_isRun = false
				logText("End Task")
				return
			default:
				if _isRun == false {
					return
				}
				if currentProcess == nil {
					if len(config.ListProcess) <= pIndex {
						pIndex = 0
					}
					currentProcess = &config.ListProcess[pIndex]
					FocusWindow(currentProcess)
					logText("Focus Window")
					pointIndex = currentProcess.PointIndex
					textIndex = currentProcess.TextIndex
					time.Sleep(1 * time.Second)
				}
				pointIndex, textIndex = doTask(
					currentProcess,
					pIndex,
					pointIndex,
					textIndex,
					isClick)
				time.Sleep(time.Duration(sleepTimeDuration) * time.Millisecond)
				cTime := makeTimestamp()
				if startTime+(int64)(currentProcess.Time) < cTime {
					pIndex += 1
					currentProcess.PointIndex = pointIndex
					currentProcess.TextIndex = textIndex
					currentProcess = nil
					startTime = cTime
					logText("Go to Next Process")
				}
				if keyTime+int64(clickTimeDuration) < cTime {
					keyTime = cTime
					isClick = true
				} else {
					isClick = false
				}
			}
		}
	}()
	logText("End Function")
	return nil
}
