package app

import(
    "fmt"
    "github.com/windwp/go-at/pkg/model"
)
// Package main provides ...
var config *model.AppConfig
func Setup() *model.AppConfig   {
    pNum:=10
    config=&model.AppConfig{
        ListProcess : make([]model.ProcessConfig, 0, pNum),
    }

    for i := 0; i < pNum; i++ {
        iP:=model.ProcessConfig{
            Pid :i,
            Name:fmt.Sprintf("process %d",i),
            Time: 10,
            Text: "test",
            Points: make([]model.Point, 0),
        }
        config.ListProcess=append(config.ListProcess, iP)
    }
    config.SelectedProcess=&config.ListProcess[0]
	return config
}
