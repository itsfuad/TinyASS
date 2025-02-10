package commands

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected Instruction
		hasError bool
	}{
		{"LOAD R1 10", Instruction{LOAD, []int{1, 10}}, false},
		{"STORE R2 0x1A", Instruction{STORE, []int{2, 26}}, false},
		{"ADD R1 R2 R3", Instruction{ADD, []int{1, 2, 3}}, false},
		{"SUB R0 R1 R2", Instruction{SUB, []int{0, 1, 2}}, false},
		{"MUL R3 R2 R1", Instruction{MUL, []int{3, 2, 1}}, false},
		{"DIV R1 R0 R3", Instruction{DIV, []int{1, 0, 3}}, false},
		{"REM R2 R1 R0", Instruction{REM, []int{2, 1, 0}}, false},
		{"AND R0 R1 R2", Instruction{AND, []int{0, 1, 2}}, false},
		{"OR R1 R2 R3", Instruction{OR, []int{1, 2, 3}}, false},
		{"XOR R2 R3 R0", Instruction{XOR, []int{2, 3, 0}}, false},
		{"NOT R3 R2", Instruction{NOT, []int{3, 2}}, false},
		{"SHL R0 R1 R2", Instruction{SHL, []int{0, 1, 2}}, false},
		{"SHR R1 R2 R3", Instruction{SHR, []int{1, 2, 3}}, false},
		{"GT R0 R1 R2", Instruction{GT, []int{0, 1, 2}}, false},
		{"LT R1 R2 R3", Instruction{LT, []int{1, 2, 3}}, false},
		{"GTE R2 R3 R0", Instruction{GTE, []int{2, 3, 0}}, false},
		{"LTE R3 R0 R1", Instruction{LTE, []int{3, 0, 1}}, false},
		{"EQ R0 R1 R2", Instruction{EQ, []int{0, 1, 2}}, false},
		{"NEQ R1 R2 R3", Instruction{NEQ, []int{1, 2, 3}}, false},
		{"JMP 0x10", Instruction{JMP, []int{16}}, false},
		{"JZ R1 0x20", Instruction{JZ, []int{1, 32}}, false},
		{"JNZ R2 0x30", Instruction{JNZ, []int{2, 48}}, false},
		{"PRINT R0", Instruction{PRINT, []int{-1, 0}}, false},
		{"PRINT MEM 0x40", Instruction{PRINT, []int{64}}, false},
		{"HALT", Instruction{HALT, []int{}}, false},
		{"INVALID", Instruction{}, true},
	}

	for _, test := range tests {
		result, err := ParseInstruction(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseInstruction(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if !compareInstructions(result, test.expected) {
			t.Errorf("ParseInstruction(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func compareInstructions(a, b Instruction) bool {
	if a.Opcode != b.Opcode {
		return false
	}
	if len(a.Operands) != len(b.Operands) {
		return false
	}
	for i := range a.Operands {
		if a.Operands[i] != b.Operands[i] {
			return false
		}
	}
	return true
}

func TestParseRegister(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"R0", 0, false},
		{"R1", 1, false},
		{"R2", 2, false},
		{"R3", 3, false},
		{"R4", 0, true},
		{"RX", 0, true},
	}

	for _, test := range tests {
		result, err := ParseRegister(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseRegister(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseRegister(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestParseMemory(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"0x00", 0, false},
		{"0x1A", 26, false},
		{"0xFF", 255, false},
		{"0x100", 0, true},
		{"0xG1", 0, true},
	}

	for _, test := range tests {
		result, err := ParseMemory(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseMemory(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseMemory(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestParseValue(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"10", 10, false},
		{"-5", -5, false},
		{"0", 0, false},
		{"abc", 0, true},
		{"", 0, true},
	}

	for _, test := range tests {
		result, err := ParseValue(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseValue(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseValue(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
