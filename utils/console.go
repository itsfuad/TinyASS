package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {

	defer func() {
		if r := recover(); r != nil {
			RED.Printf("Error in ClearScreen: %v\n", r)
		}
	}()

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
