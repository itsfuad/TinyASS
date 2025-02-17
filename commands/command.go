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
	LOAD  = iota // Load value into register
	STORE        // Store register to memory
	ADD          // Add
	SUB          // Subtract
	MUL          // Multiply
	DIV          // Divide
	REM          // Remainder
	AND          // Bitwise AND
	OR           // Bitwise OR
	XOR          // Bitwise XOR
	NOT          // Bitwise NOT
	SHL          // Shift left
	SHR          // Shift right
	GT           // Greater than
	LT           // Less than
	GTE          // Greater than or equal
	LTE          // Less than or equal
	EQ           // Equal
	NEQ          // Not equal
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
	case "AND":
		return ParseAND(parts)
	case "OR":
		return ParseOR(parts)
	case "XOR":
		return ParseXOR(parts)
	case "NOT":
		return ParseNOT(parts)
	case "SHL":
		return ParseSHL(parts)
	case "SHR":
		return ParseSHR(parts)
	case "GT":
		return ParseComparison(GT, parts)
	case "LT":
		return ParseComparison(LT, parts)
	case "GTE":
		return ParseComparison(GTE, parts)
	case "LTE":
		return ParseComparison(LTE, parts)
	case "EQ":
		return ParseComparison(EQ, parts)
	case "NEQ":
		return ParseComparison(NEQ, parts)
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

// ParseRegister validates and parses a register string formatted as "R0" to "R3".
// It checks that the input string begins with "R", is exactly 2 characters long,
// and that the following character represents a numeric value between 0 and 3 (inclusive).
// On a valid register string, the function returns the parsed register number as an int.
// If the register is invalid, it returns an error with an appropriate formatting message.
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

// ParseMemory parses a memory address provided as a hexadecimal string with a "0x" or "0X" prefix.
// It trims any surrounding whitespace, ensures the address has the correct prefix, and converts
// the hexadecimal digits into an integer. An error is returned if the address lacks the proper prefix,
// if the conversion fails, or if the resulting integer is negative or exceeds the defined limits
// (i.e., when it is not within the valid range defined by MEMORY_SIZE).
func ParseMemory(addr string) (int, error) {
	addr = strings.TrimSpace(addr)
	// if address does not start with 0x or 0X, return error
	if !(strings.HasPrefix(addr, "0x") || strings.HasPrefix(addr, "0X")) {
		return 0, fmt.Errorf(INVALID_MEMORY_ADDRESS, addr)
	}
	// convert hex string to integer. parse address from 3rd character to end
	num, err := strconv.ParseInt(addr[2:], 16, 0)
	if err != nil || int(num) < 0 || int(num) >= MEMORY_SIZE {
		return 0, fmt.Errorf(INVALID_MEMORY_ADDRESS, addr)
	}
	return int(num), nil
}

// ParseValue parses a string value into an integer. It trims any surrounding whitespace,
// and attempts to convert the string into an integer. If the conversion fails, an error
func ParseValue(val string) (int, error) {

	//validate length of value
	if len(val) == 0 {
		return 0, fmt.Errorf(INVALID_VALUE, val)
	}

	num, err := strconv.Atoi(strings.TrimSpace(val)) // convert string to integer
	if err != nil {
		return 0, fmt.Errorf(INVALID_VALUE, val)
	}
	return num, nil
}

// ParseLoad parses the LOAD instruction. It expects 3 parts: "LOAD", a register, and a value.
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

// ParseAND parses the AND instruction. It expects 4 parts: "AND", two registers, and a destination register.
func ParseAND(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("AND requires 3 operands\nExample: AND R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{AND, []int{registers[0], registers[1], registers[2]}}, nil
}

// ParseOR parses the OR instruction. It expects 4 parts: "OR", two registers, and a destination register.
func ParseOR(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("OR requires 3 operands\nExample: OR R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{OR, []int{registers[0], registers[1], registers[2]}}, nil
}

// ParseXOR parses the XOR instruction. It expects 4 parts: "XOR", two registers, and a destination register.
func ParseXOR(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("XOR requires 3 operands\nExample: XOR R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{XOR, []int{registers[0], registers[1], registers[2]}}, nil
}

// ParseNOT parses the NOT instruction. It expects 3 parts: "NOT", a register, and a destination register.
func ParseNOT(parts []string) (Instruction, error) {
	if len(parts) != 3 {
		return Instruction{}, fmt.Errorf("NOT requires 2 operands\nExample: NOT R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{NOT, []int{registers[0], registers[1]}}, nil
}

// ParseSHL parses the SHL instruction. It expects 4 parts: "SHL", two registers, and a destination register.
func ParseSHL(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("SHL requires 3 operands\nExample: SHL R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{SHL, []int{registers[0], registers[1], registers[2]}}, nil
}

// ParseSHR parses the SHR instruction. It expects 4 parts: "SHR", two registers, and a destination register.
func ParseSHR(parts []string) (Instruction, error) {
	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("SHR requires 3 operands\nExample: SHR R[0-3] R[0-3] R[0-3]")
	}
	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{SHR, []int{registers[0], registers[1], registers[2]}}, nil
}

// ParseStore parses the STORE instruction. It expects 3 parts: "STORE", a register, and a memory address.
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

// ParseRegisters parses a slice of register strings into a slice of integers.
func ParseRegisters(parts ...string) ([]int, error) {
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

// ParseAdd parses the ADD instruction. It expects 4 parts: "ADD", two registers, and a destination register.
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

// ParseSub parses the SUB instruction. It expects 4 parts: "SUB", two registers, and a destination register.
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

// ParseMul parses the MUL instruction. It expects 4 parts: "MUL", two registers, and a destination register.
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

// ParseDiv parses the DIV instruction. It expects 4 parts: "DIV", two registers, and a destination register.
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

// ParseRem parses the REM instruction. It expects 4 parts: "REM", two registers, and a destination register.
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

// ParseJmp parses the JMP instruction. It expects 2 parts: "JMP" and a memory address.
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

// ParseJz parses the JZ instruction. It expects 3 parts: "JZ", a register, and a memory address.
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

// ParseJnz parses the JNZ instruction. It expects 3 parts: "JNZ", a register, and a memory address.
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

// ParsePrint parses the PRINT instruction. It expects 2 parts: "PRINT" and a register or memory address.
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

// ParseComparison handles all comparison operations (GT, LT, GTE, LTE, EQ, NEQ)
func ParseComparison(op int, parts []string) (Instruction, error) {
	opNames := map[int]string{
		GT:  "GT",
		LT:  "LT",
		GTE: "GTE",
		LTE: "LTE",
		EQ:  "EQ",
		NEQ: "NEQ",
	}

	if len(parts) != 4 {
		return Instruction{}, fmt.Errorf("%s requires 3 operands\nExample: %s R[0-3] R[0-3] R[0-3]", opNames[op], opNames[op])
	}

	registers, err := ParseRegisters(parts[1:]...)
	if err != nil {
		return Instruction{}, err
	}

	return Instruction{op, []int{registers[0], registers[1], registers[2]}}, nil
}
