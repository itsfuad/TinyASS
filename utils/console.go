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
		if err := cmd.Run(); err != nil {
			RED.Printf("Error clearing screen: %v\n", err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			RED.Printf("Error clearing screen: %v\n", err)
		}
	}
}
