// Emulator for a tiny custom processor and its assembly language.
package main

import (
	"fmt"
	"os"
	"tinyass/runtime"
)

func main() {
	cpu := runtime.NewCPU()
	// Check if a script file or flag is passed as a command-line argument
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" {
			fmt.Println("TinyASS version 1.0.0")
			return
		}
		runtime.RunFile(cpu)
		return
	}
	// REPL mode
	runtime.StartRepl(cpu)
}
