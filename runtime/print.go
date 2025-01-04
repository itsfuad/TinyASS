package runtime

import (
	"tinyass/commands"
	"tinyass/utils"
)

func printRegisters(registers [4]int) {
	utils.BLUE.Printf("Registers: [R0=%d R1=%d R2=%d R3=%d]\n", registers[0], registers[1], registers[2], registers[3])
}

func printMemory(memory [commands.MEMORY_SIZE]int) {
	utils.BLUE.Println("Memory:")
	for i, val := range memory {
		if i == 0 {
			utils.YELLOW.Print("[ ")
		}
		//3 digit decimal number
		utils.GREY.Printf("%03d", i)
		utils.WHITE.Print(": ")
		utils.BLUE.Printf("%03d", val)
		if i == commands.MEMORY_SIZE-1 {
			utils.YELLOW.Println(" ]")
		} else {
			utils.WHITE.Print(" | ")
		}
	}
}

func printHelp() {
	utils.GREEN.Println("Commands:")
	utils.GREEN.Println("  LOAD reg val \t\t - Load value into register")
	utils.GREEN.Println("  ADD dest s1 s2 \t - Add s1 and s2 into dest")
	utils.GREEN.Println("  SUB dest s1 s2 \t - Subtract s2 from s1 into dest")
	utils.GREEN.Println("  MUL dest s1 s2 \t - Multiply s1 and s2 into dest")
	utils.GREEN.Println("  DIV dest s1 s2 \t - Divide s1 by s2 into dest")
	utils.GREEN.Println("  REM dest s1 s2 \t - Remainder of s1 divided by s2 into dest")
	utils.GREEN.Println("  JMP addr \t\t - Jump to address")
	utils.GREEN.Println("  JZ reg addr \t\t - Jump to address if register is zero")
	utils.GREEN.Println("  JNZ reg addr \t\t - Jump to address if register is not zero")
	utils.GREEN.Println("  HALT \t\t\t - Stop execution")
	utils.GREEN.Println("  reg \t\t\t - Show registers")
	utils.GREEN.Println("  mem \t\t\t - Show memory")
	utils.GREEN.Println("  PRINT Rn \t\t - Print value of register Rn")
	utils.GREEN.Println("  PRINT MEM addr \t - Print value at memory address")
	utils.GREEN.Println("  exit \t\t\t - Exit interpreter")
	utils.GREEN.Println("  cls \t\t\t - Clear the screen")
	utils.GREEN.Println("  help \t\t\t - Show this help message")
}