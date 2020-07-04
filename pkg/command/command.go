// Package command provides ...
package command

import (
	"fmt"
	"log"
	"os/exec"
    "strings"
	"github.com/windwp/go-at/pkg/model"
)

const XPROP_WINDOW="xprop -root | grep '_NET_CLIENT_LIST_STACKING(WINDOW)'"

func GetListProcess() ([]model.ProcessConfig, error) {
	lConfig := make([]model.ProcessConfig, 0)
    out, err := exec.Command("bash","-c",XPROP_WINDOW).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" %s\n", out)

    s:=string(out)
    s=strings.ReplaceAll(s,"window id #"," ")
    for _, item := range strings.Split(s,",") {
        lConfig=append(lConfig,model.ProcessConfig{
            Name: item,
        })
    }

	return lConfig, nil

}
