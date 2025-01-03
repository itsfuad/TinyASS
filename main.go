// Emulator for a tiny custom processor and its assembly language.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//self-imported package
	"tinyass/utils"
)

// Memory size
const MEMORY_SIZE = 256

// Opcodes
const (
	LOAD = iota  // Load value into register
	STORE        // Store register to memory
	ADD          // Add
	SUB          // Subtract
	MUL		  	// Multiply
	DIV		  	// Divide
	REM		  	// Remainder
	JMP          // Unconditional jump
	JZ           // Jump if zero
	JNZ          // Jump if not zero
	PRINT        // Print value
	HALT         // Stop execution
)

const INVALID_REGISTER_ERROR = "invalid register: %s\nValid registers are R0, R1, R2, R3"

// Instruction format
type Instruction struct {
	Opcode   int
	Operands []int
}

// CPU state
type CPU struct {
	memory    [MEMORY_SIZE]int
	registers [4]int      // R0-R3
	pc        int         // Program counter
	program   []Instruction
}

// Create new CPU instance
func NewCPU() *CPU {
	return &CPU{
		pc: 0,
	}
}

// Load program into memory
func (cpu *CPU) LoadProgram(instructions []Instruction) {
	cpu.program = instructions
}

// Execute one instruction
func (cpu *CPU) Execute(inst Instruction) bool {
	switch inst.Opcode {
	case LOAD:
		cpu.registers[inst.Operands[0]] = inst.Operands[1]
	case STORE:
		cpu.memory[inst.Operands[1]] = cpu.registers[inst.Operands[0]]
	case ADD:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] + cpu.registers[inst.Operands[2]]
	case SUB:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] - cpu.registers[inst.Operands[2]]
	case MUL:
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] * cpu.registers[inst.Operands[2]]
	case DIV:
		if cpu.registers[inst.Operands[2]] == 0 {
			fmt.Println("Error: Division by zero")
			return false
		}
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] / cpu.registers[inst.Operands[2]]
	case REM:
		if cpu.registers[inst.Operands[2]] == 0 {
			fmt.Println("Error: Division by zero")
			return false
		}
		cpu.registers[inst.Operands[0]] = cpu.registers[inst.Operands[1]] % cpu.registers[inst.Operands[2]]
	case JMP:
		cpu.pc = inst.Operands[0]
		return true
	case JZ:
		if cpu.registers[inst.Operands[0]] == 0 {
			cpu.pc = inst.Operands[1]
			return true
		}
	case JNZ:
		if cpu.registers[inst.Operands[0]] != 0 {
			cpu.pc = inst.Operands[1]
			return true
		}
	case PRINT:
		// Print a value from a register or memory
		if inst.Operands[0] == -1 { //
			utils.BLUE.Printf("Register R%d = %d\n", inst.Operands[1], cpu.registers[inst.Operands[1]])
		} else {
			utils.BLUE.Printf("Memory[%d] = %d\n", inst.Operands[0], cpu.memory[inst.Operands[0]])
		}
	case HALT:
		return false
	}
	return true
}


// ParseInstruction converts a string to Instruction
// ParseInstruction converts a string to Instruction
func ParseInstruction(line string) (Instruction, error) {
	// Strip comments
	commentIndex := strings.Index(line, ";")
	if commentIndex != -1 {
		line = line[:commentIndex]
	}

	line = strings.TrimSpace(line)

	// Skip empty lines
	if len(line) == 0 {
		return Instruction{}, fmt.Errorf("empty instruction")
	}

	parts := strings.Fields(line)
	opcode := parts[0]

	switch opcode {
	case "LOAD":
		return parseLoad(parts)
	case "STORE":
		return parseStore(parts)
	case "ADD":
		return parseAdd(parts)
	case "SUB":
		return parseSub(parts)
	case "MUL":
		return parseMul(parts)
	case "DIV":
		return parseDiv(parts)
	case "REM":
		return parseRem(parts)
	case "JMP":
		return parseJmp(parts)
	case "JZ":
		return parseJz(parts)
	case "JNZ":
		return parseJnz(parts)
	case "PRINT":
		return parsePrint(parts)
	case "HALT":
		return Instruction{HALT, []int{}}, nil
	default:
		return Instruction{}, fmt.Errorf("unknown instruction: %s", opcode)
	}
}


func parseLoad(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("LOAD requires 2 operands\nExample: LOAD R[0-3] val")
	}

	// Parse register and it must be a valid register
	if !isValidRegister(parts[1]) {
		return Instruction{}, fmt.Errorf(INVALID_REGISTER_ERROR, parts[1])
	}

	reg, err := strconv.Atoi(strings.TrimPrefix(parts[1], "R"))
	if err != nil || reg < 0 || reg > 3 {
		return Instruction{}, fmt.Errorf(INVALID_REGISTER_ERROR, parts[1])
	}
	val, err := strconv.Atoi(parts[2])
	if err != nil {
		return Instruction{}, fmt.Errorf("invalid value: %s", parts[2])
	}
	return Instruction{LOAD, []int{reg, val}}, nil
}

func parseStore(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("STORE requires 2 operands\nExample: STORE R[0-3] addr")
	}
	// Parse register and it must be a valid register
	if !isValidRegister(parts[1]) {
		return Instruction{}, fmt.Errorf(INVALID_REGISTER_ERROR, parts[1])
	}
	reg, err := strconv.Atoi(strings.TrimPrefix(parts[1], "R"))
	if err != nil || reg < 0 || reg > 3 {
		return Instruction{}, fmt.Errorf(INVALID_REGISTER_ERROR, parts[1])
	}
	addr, err := strconv.Atoi(parts[2])
	if err != nil {
		return Instruction{}, fmt.Errorf("invalid address: %s", parts[2])
	}
	return Instruction{STORE, []int{reg, addr}}, nil
}

func isValidRegister(reg string) bool {
	return strings.HasPrefix(reg, "R") && len(reg) == 2 && reg[1] >= '0' && reg[1] <= '3'
}

func parseAdd(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("ADD requires 3 operands\nExample: ADD R[0-3] R[0-3] R[0-3]")
	}
	dest, _ := strconv.Atoi(parts[1])
	src1, _ := strconv.Atoi(parts[2])
	src2, _ := strconv.Atoi(parts[3])
	return Instruction{ADD, []int{dest, src1, src2}}, nil
}

func parseSub(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("SUB requires 3 operands\nExample: SUB R[0-3] R[0-3] R[0-3]")
	}
	dest, _ := strconv.Atoi(parts[1])
	src1, _ := strconv.Atoi(parts[2])
	src2, _ := strconv.Atoi(parts[3])
	return Instruction{SUB, []int{dest, src1, src2}}, nil
}

func parseMul(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("MUL requires 3 operands\nExample: MUL R[0-3] R[0-3] R[0-3]")
	}
	dest, _ := strconv.Atoi(parts[1])
	src1, _ := strconv.Atoi(parts[2])
	src2, _ := strconv.Atoi(parts[3])
	return Instruction{MUL, []int{dest, src1, src2}}, nil
}

func parseDiv(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("DIV requires 3 operands\nExample: DIV R[0-3] R[0-3] R[0-3]")
	}
	dest, _ := strconv.Atoi(parts[1])
	src1, _ := strconv.Atoi(parts[2])
	src2, _ := strconv.Atoi(parts[3])
	return Instruction{DIV, []int{dest, src1, src2}}, nil
}

func parseRem(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("REM requires 3 operands\nExample: REM R[0-3] R[0-3] R[0-3]")
	}
	dest, _ := strconv.Atoi(parts[1])
	src1, _ := strconv.Atoi(parts[2])
	src2, _ := strconv.Atoi(parts[3])
	return Instruction{REM, []int{dest, src1, src2}}, nil
}

func parseJmp(parts []string) (Instruction, error) {
	if len(parts) != 2 {
		return Instruction{}, fmt.Errorf("JMP requires 1 operand\nExample: JMP addr")
	}
	addr, _ := strconv.Atoi(parts[1])
	return Instruction{JMP, []int{addr}}, nil
}

func parseJz(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("JZ requires 2 operands\nExample: JZ R[0-3] addr")
	}
	reg, _ := strconv.Atoi(parts[1])
	addr, _ := strconv.Atoi(parts[2])
	return Instruction{JZ, []int{reg, addr}}, nil
}

func parseJnz(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("JNZ requires 2 operands\nExample: JNZ R[0-3] addr")
	}
	reg, _ := strconv.Atoi(parts[1])
	addr, _ := strconv.Atoi(parts[2])
	return Instruction{JNZ, []int{reg, addr}}, nil
}

func parsePrint(parts []string) (Instruction, error) {
    if len(parts) != 2 && len(parts) != 3 {
        return Instruction{}, fmt.Errorf("PRINT requires 1 or 2 operands\nExample: PRINT R0 or PRINT MEM 100")
    }

    if parts[1] == "MEM" {
        // Print memory value
		if len(parts) != 3 {
			return Instruction{}, fmt.Errorf("PRINT MEM requires 1 operand\nExample: PRINT MEM 100")
		}
        addr, err := strconv.Atoi(parts[2])
        if err != nil {
            return Instruction{}, fmt.Errorf("invalid memory address: %s", parts[2])
        }
        return Instruction{PRINT, []int{addr}}, nil
    } else if strings.HasPrefix(parts[1], "R") {
        // Print register value
        reg, err := strconv.Atoi(strings.TrimPrefix(parts[1], "R"))
        if err != nil || reg < 0 || reg > 3 {
            return Instruction{}, fmt.Errorf("invalid register: %s", parts[1])
        }
        return Instruction{PRINT, []int{-1, reg}}, nil
    }

    return Instruction{}, fmt.Errorf("invalid operand for PRINT")
}


func main() {
	cpu := NewCPU()

	// Check if a script file is passed as a command-line argument
	if len(os.Args) > 1 {
		// Load program from file
		runFile(cpu)
		return
	}

	// REPL mode
	startRepl(cpu)
}

func runFile(cpu *CPU) {

	filename := os.Args[1]
	// Read and execute the script
	script, err := os.ReadFile(filename)
	if err != nil {
		utils.RED.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	// Parse the script into instructions
	lines := strings.Split(string(script), "\n")
	var instructions []Instruction
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") { // Skip empty lines or comments
			continue
		}

		inst, err := ParseInstruction(line)
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

func startRepl(cpu *CPU) {
	scanner := bufio.NewScanner(os.Stdin)
	utils.GREEN.Println("Tiny Assembly Interpreter")
	utils.BLUE.Println("Type 'help' for commands, 'exit' to quit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		switch strings.ToLower(line) {
		case "exit":
			return
		case "help":
			printHelp()
			continue
		case "reg":
			printRegisters(cpu.registers)
			continue
		case "mem":
			printMemory(cpu.memory)
			continue
		}

		inst, err := ParseInstruction(strings.ToUpper(line))
		if err != nil {
			utils.RED.Printf("Error: %v\n", err)
			continue
		}

		cpu.pc++
		cpu.Execute(inst)
	}
}

func printRegisters(registers [4]int) {
	utils.BLUE.Printf("Registers: [R0=%d R1=%d R2=%d R3=%d]\n", registers[0], registers[1], registers[2], registers[3])
}

func printMemory(memory [MEMORY_SIZE]int) {
	utils.BLUE.Println("Memory:")
	for i, val := range memory {
		if i == 0 {
			utils.YELLOW.Print("[ ")
		}
		utils.GREY.Printf("0x%02X", i)
		utils.WHITE.Print(": ")
		utils.BLUE.Print(val)
		if i == MEMORY_SIZE-1 {
			utils.YELLOW.Println(" ]")
		} else {
			utils.WHITE.Print(", ")
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
}