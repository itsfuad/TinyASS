package runtime

import (
	"bufio"
	"os"
	"strings"

	"tinyass/commands"
	"tinyass/utils"
)

// State encapsulates the CPU's registers and memory.
type State struct {
	memory    [commands.MEMORY_SIZE]int
	registers [4]int // R0-R3
}

// CPU state
type CPU struct {
	state    State
	pc       int // Program counter
	program  []commands.Instruction
	executor Executor
}

// Executor interface for executing instructions.
type Executor interface {
	Execute(cpu *CPU, inst commands.Instruction) bool
}

// CPUExecutor is a concrete implementation of the Executor interface.
type CPUExecutor struct{}

// NewCPUExecutor creates a new CPUExecutor instance.
func NewCPUExecutor() *CPUExecutor {
	return &CPUExecutor{}
}

// Create new CPU instance
func NewCPU() *CPU {
	cpu := &CPU{
		pc:       0,
		executor: NewCPUExecutor(),
	}
	return cpu
}

// Load program into memory
func (cpu *CPU) LoadProgram(instructions []commands.Instruction) {
	cpu.program = instructions
}

// ExecuteOne executes a single instruction.
func (cpu *CPU) ExecuteOne() bool {
	if cpu.pc >= len(cpu.program) {
		return false
	}
	inst := cpu.program[cpu.pc]
	cpu.pc++
	return cpu.executor.Execute(cpu, inst)
}

// Execute implements the Executor interface for CPUExecutor.
func (e *CPUExecutor) Execute(cpu *CPU, inst commands.Instruction) bool {
	switch inst.Opcode {
	case commands.LOAD:
		e.executeLoad(cpu, inst)
	case commands.STORE:
		e.executeStore(cpu, inst)
	case commands.ADD:
		e.executeAdd(cpu, inst)
	case commands.SUB:
		e.executeSub(cpu, inst)
	case commands.MUL:
		e.executeMul(cpu, inst)
	case commands.DIV:
		return e.executeDiv(cpu, inst)
	case commands.REM:
		return e.executeRem(cpu, inst)
	case commands.AND:
		e.executeAnd(cpu, inst)
	case commands.OR:
		e.executeOr(cpu, inst)
	case commands.XOR:
		e.executeXor(cpu, inst)
	case commands.NOT:
		e.executeNot(cpu, inst)
	case commands.SHL:
		e.executeShl(cpu, inst)
	case commands.SHR:
		e.executeShr(cpu, inst)
	case commands.GT:
		e.executeGt(cpu, inst)
	case commands.LT:
		e.executeLt(cpu, inst)
	case commands.GTE:
		e.executeGte(cpu, inst)
	case commands.LTE:
		e.executeLte(cpu, inst)
	case commands.EQ:
		e.executeEq(cpu, inst)
	case commands.NEQ:
		e.executeNeq(cpu, inst)
	case commands.JMP:
		return e.executeJmp(cpu, inst)
	case commands.JZ:
		return e.executeJz(cpu, inst)
	case commands.JNZ:
		return e.executeJnz(cpu, inst)
	case commands.PRINT:
		e.executePrint(cpu, inst)
	case commands.HALT:
		return false
	}
	return true
}

func (e *CPUExecutor) executeLoad(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = inst.Operands[1]
}

func (e *CPUExecutor) executeStore(cpu *CPU, inst commands.Instruction) {
	cpu.state.memory[inst.Operands[1]] = cpu.state.registers[inst.Operands[0]]
}

func (e *CPUExecutor) executeAdd(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] + cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeSub(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] - cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeMul(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] * cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeDiv(cpu *CPU, inst commands.Instruction) bool {
	if cpu.state.registers[inst.Operands[2]] == 0 {
		utils.RED.Printf("Error: Division by zero on program counter %d\n", cpu.pc)
		return false
	}
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] / cpu.state.registers[inst.Operands[2]]
	return true
}

func (e *CPUExecutor) executeRem(cpu *CPU, inst commands.Instruction) bool {
	if cpu.state.registers[inst.Operands[2]] == 0 {
		utils.RED.Println("Error: Division by zero")
		return false
	}
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] % cpu.state.registers[inst.Operands[2]]
	return true
}

func (e *CPUExecutor) executeAnd(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] & cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeOr(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] | cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeXor(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] ^ cpu.state.registers[inst.Operands[2]]
}

func (e *CPUExecutor) executeNot(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = ^cpu.state.registers[inst.Operands[1]]
}

func (e *CPUExecutor) executeShl(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] << uint(inst.Operands[2])
}

func (e *CPUExecutor) executeShr(cpu *CPU, inst commands.Instruction) {
	cpu.state.registers[inst.Operands[0]] = cpu.state.registers[inst.Operands[1]] >> uint(inst.Operands[2])
}

func (e *CPUExecutor) executeGt(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] > cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeLt(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] < cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeGte(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] >= cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeLte(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] <= cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeEq(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] == cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeNeq(cpu *CPU, inst commands.Instruction) {
	if cpu.state.registers[inst.Operands[1]] != cpu.state.registers[inst.Operands[2]] {
		cpu.state.registers[inst.Operands[0]] = 1
	} else {
		cpu.state.registers[inst.Operands[0]] = 0
	}
}

func (e *CPUExecutor) executeJmp(cpu *CPU, inst commands.Instruction) bool {
	cpu.pc = inst.Operands[0]
	return true
}

func (e *CPUExecutor) executeJz(cpu *CPU, inst commands.Instruction) bool {
	if cpu.state.registers[inst.Operands[0]] == 0 {
		cpu.pc = inst.Operands[1]
		return true
	}
	return false
}

func (e *CPUExecutor) executeJnz(cpu *CPU, inst commands.Instruction) bool {
	if cpu.state.registers[inst.Operands[0]] != 0 {
		cpu.pc = inst.Operands[1]
		return true
	}
	return false
}

func (e *CPUExecutor) executePrint(cpu *CPU, inst commands.Instruction) {
	if inst.Operands[0] == -1 {
		utils.BLUE.Printf("%d\n", cpu.state.registers[inst.Operands[1]])
	} else {
		utils.BLUE.Printf("%d\n", cpu.state.memory[inst.Operands[0]])
	}
}

// RunFile reads a file, parses the instructions, loads them into the CPU, and executes them.
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
		if !cpu.ExecuteOne() {
			break
		}
	}

	utils.GREEN.Println("Execution completed.")
}

// StartRepl starts the interactive REPL.
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
			printRegisters(cpu.state.registers)
			continue
		case "mem":
			printMemory(cpu.state.memory)
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

		cpu.program = []commands.Instruction{inst}
		cpu.pc = 0
		cpu.ExecuteOne()
	}
}
