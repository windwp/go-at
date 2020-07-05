// Package command provides ...
package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/windwp/go-at/pkg/model"
)

const XPROP_WINDOW_CMD = "xprop -root | grep '_NET_CLIENT_LIST_STACKING(WINDOW)'"
const PROCESS_INFO_CMD = "xprop -id %s"

func GetListProcess() ([]model.ProcessConfig, error) {
	lConfig := make([]model.ProcessConfig, 0)
	out, err := exec.Command("bash", "-c", XPROP_WINDOW_CMD).Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.ReplaceAll(s, "_NET_CLIENT_LIST_STACKING(WINDOW): window id # ", " ")
	for _, item := range strings.Split(s, ",") {
		config, e := getProcessInfomation(item)
		if e == nil {
			lConfig = append(lConfig, *config)
		}
	}
	return lConfig, nil

}

func getProcessInfomation(windowid string) (*model.ProcessConfig, error) {
	result := model.ProcessConfig{}
	out, err := exec.Command("bash", "-c", fmt.Sprintf(PROCESS_INFO_CMD, windowid)).Output()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	query := string(out)
	lines := strings.Split(query, "\n")
	for _, item := range lines {
		params := strings.Split(item, "=")
		if len(params) == 2 {
			fParam := strings.Trim(params[0], " ")
			switch fParam {
			case "_NET_WM_NAME(UTF8_STRING)":
				result.Title = strings.ReplaceAll(params[1], "\"", "")
				break
			case "_NET_WM_PID(CARDINAL)":
				result.Pid, _ = strconv.Atoi(strings.Trim(params[1], " "))
				break
			}
		}
	}
	return &result, nil

}

func SaveJSON(config *model.AppConfig) error {
	configDir, _ := os.UserHomeDir()
	path := configDir + "/" + model.DATA_PATH
	result, err := json.Marshal(config)
	if err == nil {
        err:=ioutil.WriteFile(path, result, 0644)
        if err != nil {
            return err
        }
	} else {
		return err
	}
	return nil
}

func LoadJson() (*model.AppConfig, error) {
	configDir, _ := os.UserHomeDir()
	path := configDir + "/" + model.DATA_PATH
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil
	}
	var config model.AppConfig
	err = json.Unmarshal(jsonData, &config)

	if err == nil {
		return &config, nil
	}
    os.Remove(path)
	return nil, errors.New("Json error")
}
