// Emulator for a tiny custom processor and its assembly language.
package main

import (
	"os"
	"tinyass/runtime"
)


func main() {
	cpu := runtime.NewCPU()
	// Check if a script file is passed as a command-line argument
	if len(os.Args) > 1 {
		// Load program from file
		runtime.RunFile(cpu)
		return
	}
	// REPL mode
	runtime.StartRepl(cpu)
}

