package runtime

import (
	"tinyass/commands"
	"tinyass/utils"
)

func printRegisters(registers [4]int) {
	// Improved register display with header and different colors for 0 and nonzero values
	utils.BLUE.Println("---- REGISTERS ----")
	// Use GREY for 0 and GREEN for nonzero
	regStr := ""
	for i, val := range registers {
		if val == 0 {
			regStr += utils.GREY.Sprintf("R%d: %d  ", i, val)
		} else {
			regStr += utils.GREEN.Sprintf("R%d: %d  ", i, val)
		}
	}
	utils.GREEN.Println(regStr)
	utils.BLUE.Println("-------------------")
}

func printMemory(memory [commands.MEMORY_SIZE]int) {
	// Improved memory display in 16-value rows with header and different colors for 0 and nonzero values
	utils.BLUE.Println("---- MEMORY ----")
	for i := 0; i < commands.MEMORY_SIZE; i++ {
		var cell string
		if memory[i] == 0 {
			cell = utils.GREY.Sprintf("%02X: %03d  ", i, memory[i])
		} else {
			cell = utils.BOLD_YELLOW.Sprintf("%02X: %03d  ", i, memory[i])
		}
		utils.GREY.Print(cell)
		if (i+1)%16 == 0 {
			utils.YELLOW.Println("")
		}
	}
	utils.BLUE.Println("----------------")
}

func printHelp() {
	utils.GREEN.Println("Commands:")
	utils.GREEN.Println("  LOAD reg val      \t - Load value into register")
	utils.GREEN.Println("  STORE reg addr    \t - Store value from register into memory address")
	utils.GREEN.Println("  ADD dest s1 s2    \t - Add s1 and s2 into dest")
	utils.GREEN.Println("  SUB dest s1 s2    \t - Subtract s2 from s1 into dest")
	utils.GREEN.Println("  MUL dest s1 s2    \t - Multiply s1 and s2 into dest")
	utils.GREEN.Println("  DIV dest s1 s2    \t - Divide s1 by s2 into dest")
	utils.GREEN.Println("  REM dest s1 s2    \t - Remainder of s1 divided by s2 into dest")
	utils.GREEN.Println("  AND dest s1 s2    \t - Bitwise AND of s1 and s2 into dest")
	utils.GREEN.Println("  OR dest s1 s2     \t - Bitwise OR of s1 and s2 into dest")
	utils.GREEN.Println("  XOR dest s1 s2    \t - Bitwise XOR of s1 and s2 into dest")
	utils.GREEN.Println("  NOT dest s1       \t - Bitwise NOT of s1 into dest")
	utils.GREEN.Println("  SHL dest s1 s2    \t - Shift s1 left by s2 bits into dest")
	utils.GREEN.Println("  SHR dest s1 s2    \t - Shift s1 right by s2 bits into dest")
	utils.GREEN.Println("  GT dest s1 s2     \t - Set dest to 1 if s1 > s2, else 0")
	utils.GREEN.Println("  LT dest s1 s2     \t - Set dest to 1 if s1 < s2, else 0")
	utils.GREEN.Println("  GTE dest s1 s2    \t - Set dest to 1 if s1 >= s2, else 0")
	utils.GREEN.Println("  LTE dest s1 s2    \t - Set dest to 1 if s1 <= s2, else 0")
	utils.GREEN.Println("  EQ dest s1 s2     \t - Set dest to 1 if s1 == s2, else 0")
	utils.GREEN.Println("  NEQ dest s1 s2    \t - Set dest to 1 if s1 != s2, else 0")
	utils.GREEN.Println("  JMP addr          \t - Jump to address")
	utils.GREEN.Println("  JZ reg addr       \t - Jump to address if register is zero")
	utils.GREEN.Println("  JNZ reg addr      \t - Jump to address if register is not zero")
	utils.GREEN.Println("  HALT              \t - Stop execution")
	utils.GREEN.Println("  reg               \t - Show registers")
	utils.GREEN.Println("  mem               \t - Show memory")
	utils.GREEN.Println("  PRINT Rn          \t - Print value of register Rn")
	utils.GREEN.Println("  PRINT MEM addr    \t - Print value at memory address")
	utils.GREEN.Println("  version           \t - Show version info")
	utils.GREEN.Println("  exit              \t - Exit interpreter")
	utils.GREEN.Println("  cls               \t - Clear the screen")
	utils.GREEN.Println("  help              \t - Show this help message")
}
