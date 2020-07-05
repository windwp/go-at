package app

import (
	"log"

	"github.com/thoas/go-funk"
	"github.com/windwp/go-at/pkg/command"
	"github.com/windwp/go-at/pkg/model"
)

// Package main provides ...
var config *model.AppConfig

func Setup() *model.AppConfig {
	pNum := 10
	config = &model.AppConfig{
		Status:      "Idle",
		Message:     "OK",
		ListProcess: make([]model.ProcessConfig, 0, pNum),
	}
	jsonConfig, err := command.LoadJson()
	if err == nil &&jsonConfig !=nil {
		config = jsonConfig
	}

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
		config.ListProcess = funk.Filter(
			config.ListProcess,
			func(item model.ProcessConfig) bool {
				return item.Name != config.SelectedProcess.Name
			},
		).([]model.ProcessConfig)
		config.SelectedProcess = nil
	}
	return 0
}
