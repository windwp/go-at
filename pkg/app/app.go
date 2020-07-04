package app

import(
    "fmt"
)
// Package main provides ...
type Point struct {
	x, y int
}
type ProcessConfig struct {
	Pid       int
	Name      string
	Time   int
	Text      string
	Points []Point
}

type AppConfig struct{
	ListProcess []ProcessConfig
}

var config AppConfig
func Setup() AppConfig   {
    pNum:=10
    config=AppConfig{
        ListProcess : make([]ProcessConfig, 0, pNum),
    }
    for i := 0; i < pNum; i++ {
        iP:=ProcessConfig{
            Pid :0,
            Name:fmt.Sprintf("process %d",i),
            Time: 10,
            Text: "test",
            Points: make([]Point, 0),
        }
        config.ListProcess=append(config.ListProcess, iP)
    }
	return config
}
