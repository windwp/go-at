// Package testing provides ...
package test

import (
	"fmt"
	"testing"

	"github.com/windwp/go-at/pkg/app"
	"github.com/windwp/go-at/pkg/command"
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
	config,err := command.LoadJson()
	if err != nil {
		t.Errorf("load json error %v", err)
	}
    if(len(config.ListProcess)<1){
        t.Error("load list process error")
    }
    fmt.Printf("Message %s\n",config.Message)

}
