package commands
import (
	"fmt"
	"strconv"
	"strings"
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
const INVALID_MEMORY_ADDRESS = "invalid memory address: %s\nValid memory addresses are 0x00 to 0xFF"
const INVALID_VALUE = "invalid value: %s"

// Instruction format
type Instruction struct {
	Opcode   int
	Operands []int
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
		return ParseLoad(parts)
	case "STORE":
		return ParseStore(parts)
	case "ADD":
		return ParseAdd(parts)
	case "SUB":
		return ParseSub(parts)
	case "MUL":
		return ParseMul(parts)
	case "DIV":
		return ParseDiv(parts)
	case "REM":
		return ParseRem(parts)
	case "JMP":
		return ParseJmp(parts)
	case "JZ":
		return ParseJz(parts)
	case "JNZ":
		return ParseJnz(parts)
	case "PRINT":
		return ParsePrint(parts)
	case "HALT":
		return Instruction{HALT, []int{}}, nil
	default:
		return Instruction{}, fmt.Errorf("unknown instruction: %s", opcode)
	}
}

func ParseRegister(reg string) (int, error) {
	if !(strings.HasPrefix(reg, "R") && len(reg) == 2 && reg[1] >= '0' && reg[1] <= '3') {
		return 0, fmt.Errorf(INVALID_REGISTER_ERROR, reg)
	}
	num, err := strconv.Atoi(strings.TrimPrefix(reg, "R"))
	if err != nil || num < 0 || num > 3 {
		return 0, fmt.Errorf(INVALID_REGISTER_ERROR, reg)
	}
	return num, nil
}

func ParseMemory(addr string) (int, error) {
	num, err := strconv.Atoi(strings.TrimSpace(addr))
	if err != nil || num < 0 || num >= MEMORY_SIZE {
		return 0, fmt.Errorf(INVALID_MEMORY_ADDRESS, addr)
	}
	return num, nil
}

func ParseValue(val string) (int, error) {

	//validate length of value
	if len(val) == 0 {
		return 0, fmt.Errorf(INVALID_VALUE, val)
	}

	num, err := strconv.Atoi(strings.TrimSpace(val))
	if err != nil {
		return 0, fmt.Errorf(INVALID_VALUE, val)
	}
	return num, nil
}

func ParseLoad(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("LOAD requires 2 operands\nExample: LOAD R[0-3] val")
	}

	reg, err := ParseRegister(parts[1])
	if err != nil {
		return Instruction{}, err
	}

	val, err := ParseValue(parts[2])
	if err != nil {
		return Instruction{}, err
	}

	return Instruction{LOAD, []int{reg, val}}, nil
}

func ParseStore(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("STORE requires 2 operands\nExample: STORE R[0-3] addr")
	}

	reg, err := ParseRegister(parts[1])
	if err != nil {
		return Instruction{}, err
	}

	addr, err := ParseMemory(parts[2])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{STORE, []int{reg, addr}}, nil
}

func ParseRegisters(parts... string) ([]int, error) {
	var regs []int
	for _, reg := range parts {
		num, err := ParseRegister(reg)
		if err != nil {
			return nil, err
		}
		regs = append(regs, num)
	}
	return regs, nil
}

func ParseAdd(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("ADD requires 3 operands\nExample: ADD R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{ADD, []int{registers[0], registers[1], registers[2]}}, nil
}

func ParseSub(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("SUB requires 3 operands\nExample: SUB R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{SUB, []int{registers[0], registers[1], registers[2]}}, nil
}

func ParseMul(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("MUL requires 3 operands\nExample: MUL R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{MUL, []int{registers[0], registers[1], registers[2]}}, nil
}

func ParseDiv(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("DIV requires 3 operands\nExample: DIV R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{DIV, []int{registers[0], registers[1], registers[2]}}, nil
}

func ParseRem(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("REM requires 3 operands\nExample: REM R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{REM, []int{registers[0], registers[1], registers[2]}}, nil
}

func ParseJmp(parts []string) (Instruction, error) {
	if len(parts) != 2 {
		return Instruction{}, fmt.Errorf("JMP requires 1 operand\nExample: JMP addr")
	}
	addr, err := ParseMemory(parts[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{JMP, []int{addr}}, nil
}

func ParseJz(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("JZ requires 2 operands\nExample: JZ R[0-3] addr")
	}
	reg, err := ParseRegister(parts[1])
	if err != nil {
		return Instruction{}, err
	}
	addr, err := ParseMemory(parts[2])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{JZ, []int{reg, addr}}, nil
}

func ParseJnz(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("JNZ requires 2 operands\nExample: JNZ R[0-3] addr")
	}
	reg, err := ParseRegister(parts[1])
	if err != nil {
		return Instruction{}, err
	}
	addr, err := ParseMemory(parts[2])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{JNZ, []int{reg, addr}}, nil
}

func ParsePrint(parts []string) (Instruction, error) {
    if len(parts) != 2 && len(parts) != 3 {
        return Instruction{}, fmt.Errorf("PRINT requires 1 or 2 operands\nExample: PRINT R0 or PRINT MEM 100")
    }

    if parts[1] == "MEM" {
        // Print memory value
		if len(parts) != 3 {
			return Instruction{}, fmt.Errorf("PRINT MEM requires 1 operand\nExample: PRINT MEM 100")
		}
        addr, err := ParseMemory(parts[2])
		if err != nil {
			return Instruction{}, err
		}
        return Instruction{PRINT, []int{addr}}, nil
    } else if strings.HasPrefix(parts[1], "R") {
        // Print register value
        reg, err := ParseRegister(parts[1])
		if err != nil {
			return Instruction{}, err
		}
        return Instruction{PRINT, []int{-1, reg}}, nil
    }

    return Instruction{}, fmt.Errorf("invalid operand for PRINT")
}