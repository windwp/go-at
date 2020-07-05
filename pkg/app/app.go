package app

import (
	"fmt"
	"log"

	"github.com/windwp/go-at/pkg/model"
)

// Package main provides ...
var config *model.AppConfig

func Setup() *model.AppConfig {
	pNum := 10
	config = &model.AppConfig{
		Status: "Idle",

		Message:     "OK",
		ListProcess: make([]model.ProcessConfig, 0, pNum),
	}

	for i := 0; i < pNum; i++ {
		iP := model.ProcessConfig{
			Pid:    i,
			Name:   fmt.Sprintf("process %d", i),
			Time:   10,
			Text:   "test",
			Points: make([]model.Point, 0),
		}
		config.ListProcess = append(config.ListProcess, iP)
	}
	config.SelectedProcess = &config.ListProcess[0]
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
