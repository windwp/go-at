// Package testing provides ...
package test

import (
	"fmt"
	"testing"

	"github.com/windwp/go-at/pkg/command"
)

func TestListProcess(t *testing.T) {
	lProcess, _ := command.GetListProcess()
	fmt.Printf("Process Length %d", len(lProcess))
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
