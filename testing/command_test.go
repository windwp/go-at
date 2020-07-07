// Package testing provides ...
package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/windwp/go-at/pkg/app"
	"github.com/windwp/go-at/pkg/command"
	"github.com/windwp/go-at/pkg/model"
)

func TestListProcess(t *testing.T) {
	lProcess, _ := command.GetListProcess()
	if len(lProcess) < 1 {
		t.Errorf("Process Empty %d", len(lProcess))
	}
	if len(lProcess) > 0 {
		fp := lProcess[0]
		fmt.Print(fp)
		if fp.Pid == 0 {
			t.Error("pid is zero")
		}
	}
}

func TestSaveJson(t *testing.T) {
	config := app.Setup()
	err := command.SaveJSON(config)
	if err != nil {
		t.Errorf("Can't save json %v", err)
	}

}

func TestLoadJson(t *testing.T) {
	config, err := command.LoadJson()
	if err != nil {
		t.Errorf("load json error %v", err)
	}
	if len(config.ListProcess) < 1 {
		t.Error("load list process error")
	}
	fmt.Printf("Message %s\n", config.Message)

}

func TestStartOneTaskTerminal(t *testing.T) {
	config := app.Setup()
	config.ListProcess = make([]model.ProcessConfig, 0)
	psList, _ := command.GetListProcess()
	for _, p := range psList {
		if strings.Index(p.Title, "@") != -1 {
			p.Text = "Hellow world\n Newline\n ok"
			p.Time = 240
			p.Points = []model.Point{
				{X: 527, Y: 244},
				{X: 923, Y: 221},
				{X: 1611, Y: 376},
			}
			config.ListProcess = append(config.ListProcess, p)
		}
	}
	if len(config.ListProcess) < 1 {
		t.Error("List Process empty")
	}
	fmt.Println("Start Task")
	command.StartTask(config, true)
	<-command.WaitTask()
}

func TestStartTaskConfig(t *testing.T) {
	config := app.Setup()
	fmt.Println("Start Task")
	fmt.Printf("Task Length %d", len(config.ListProcess))
	err := command.StartTask(config, true)
	if err != nil {
		t.Errorf(" Can't start %s", err.Error())
	}
	<-command.WaitTask()
}
