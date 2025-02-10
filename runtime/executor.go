package runtime

import (
	"bufio"
	"os"
	"strings"

	"tinyass/commands"
	"tinyass/utils"
)

// CPU state
type CPU struct {
	memory    [commands.MEMORY_SIZE]int
	registers [4]int // R0-R3
	pc        int    // Program counter
	program   []commands.Instruction
}

// Create new CPU instance
func NewCPU() *CPU {
	return &CPU{
		pc: 0,
	}
}

// Load program into memory
func (cpu *CPU) LoadProgram(instructions []commands.Instruction) {
	cpu.program = instructions
}

// Execute one instruction
func (cpu *CPU) Execute(inst commands.Instruction) bool {
	switch inst.Opcode {
	case commands.LOAD:
		cpu.registers[inst.Operands[0]] = inst.Operands[1]
	case commands.STORE:
		cpu.memory[inst.Operands[1]] = cpu.registers[inst.Operands[0]]
	case commands.ADD:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] + cpu.registers[inst.Operands[2]]
	case commands.SUB:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] - cpu.registers[inst.Operands[2]]
	case commands.MUL:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] * cpu.registers[inst.Operands[2]]
	case commands.DIV:
		if cpu.registers[inst.Operands[2]] == 0 {
			utils.RED.Printf("Error: Division by zero on program counter %d\n", cpu.pc)
			return false
		}
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] / cpu.registers[inst.Operands[2]]
	case commands.REM:
		if cpu.registers[inst.Operands[2]] == 0 {
			utils.RED.Println("Error: Division by zero")
			return false
		}
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] % cpu.registers[inst.Operands[2]]
	case commands.AND:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] & cpu.registers[inst.Operands[2]]
	case commands.OR:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] | cpu.registers[inst.Operands[2]]
	case commands.XOR:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] ^ cpu.registers[inst.Operands[2]]
	case commands.NOT:
		cpu.registers[inst.Operands[0]] = ^cpu.registers[inst.Operands[1]]
	case commands.SHL:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] << uint(inst.Operands[2])
	case commands.SHR:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] >> uint(inst.Operands[2])
	case commands.JMP:
		cpu.pc = inst.Operands[0]
		return true
	case commands.JZ:
		if cpu.registers[inst.Operands[0]] == 0 {
			cpu.pc = inst.Operands[1]
			return true
		}
	case commands.JNZ:
		if cpu.registers[inst.Operands[0]] != 0 {
			cpu.pc = inst.Operands[1]
			return true
		}
	case commands.PRINT:
		// Print a value from a register or memory
		if inst.Operands[0] == -1 { //
			utils.BLUE.Printf("Register R%d = %d\n", inst.Operands[1], cpu.registers[inst.Operands[1]])
		} else {
			utils.BLUE.Printf("Memory[%d] = %d\n", inst.Operands[0], cpu.memory[inst.Operands[0]])
		}
	case commands.HALT:
		return false
	}
	return true
}

func RunFile(cpu *CPU) {

	filename := os.Args[1]
	// Read and execute the script
	script, err := os.ReadFile(filename)
	if err != nil {
		utils.RED.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	// Parse the script into instructions
	lines := strings.Split(string(script), "\n")
	var instructions []commands.Instruction
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") { // Skip empty lines or comments
			continue
		}

		inst, err := commands.ParseInstruction(line)
		if err != nil {
			utils.RED.Printf("Error parsing line %d: %v\n", i+1, err)
			return
		}
		instructions = append(instructions, inst)
	}

	// Load the instructions into the CPU and execute them
	cpu.LoadProgram(instructions)
	for cpu.pc < len(cpu.program) {
		inst := cpu.program[cpu.pc]
		cpu.pc++
		if !cpu.Execute(inst) {
			break
		}
	}

	utils.GREEN.Println("Execution completed.")
}

func StartRepl(cpu *CPU) {
	scanner := bufio.NewScanner(os.Stdin)
	utils.GREEN.Println("Tiny Assembly Interpreter")
	utils.BLUE.Println("Type 'help' for commands, 'exit' to quit")

	for {
		utils.BLUE.Print("TinyASS > ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		switch strings.ToLower(line) {
		case "exit":
			return
		case "cls":
			utils.ClearScreen()
			continue
		case "help":
			printHelp()
			continue
		case "reg":
			printRegisters(cpu.registers)
			continue
		case "mem":
			printMemory(cpu.memory)
			continue
		case "version":
			utils.GREEN.Println("TinyASS version 1.0.0")
			continue
		}

		inst, err := commands.ParseInstruction(strings.ToUpper(line))
		if err != nil {
			utils.RED.Printf("Error: %v\n", err)
			continue
		}

		cpu.pc++
		cpu.Execute(inst)
	}
}
