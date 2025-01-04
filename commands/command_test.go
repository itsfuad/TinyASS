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
		{"LOAD R0 10", Instruction{LOAD, []int{0, 10}}, false},
		{"STORE R1 20", Instruction{STORE, []int{1, 20}}, false},
		{"ADD R0 R1 R2", Instruction{ADD, []int{0, 1, 2}}, false},
		{"SUB R2 R1 R0", Instruction{SUB, []int{2, 1, 0}}, false},
		{"MUL R0 R1 R2", Instruction{MUL, []int{0, 1, 2}}, false},
		{"DIV R2 R1 R0", Instruction{DIV, []int{2, 1, 0}}, false},
		{"REM R0 R1 R2", Instruction{REM, []int{0, 1, 2}}, false},
		{"AND R0 R1 R2", Instruction{AND, []int{0, 1, 2}}, false},
		{"OR R2 R1 R0", Instruction{OR, []int{2, 1, 0}}, false},
		{"XOR R0 R1 R2", Instruction{XOR, []int{0, 1, 2}}, false},
		{"NOT R0 R1", Instruction{NOT, []int{0, 1}}, false},
		{"SHL R0 R1 R2", Instruction{SHL, []int{0, 1, 2}}, false},
		{"SHR R2 R1 R0", Instruction{SHR, []int{2, 1, 0}}, false},
		{"JMP 100", Instruction{JMP, []int{100}}, false},
		{"JZ R0 100", Instruction{JZ, []int{0, 100}}, false},
		{"JNZ R1 200", Instruction{JNZ, []int{1, 200}}, false},
		{"PRINT R0", Instruction{PRINT, []int{-1, 0}}, false},
		{"PRINT MEM 100", Instruction{PRINT, []int{100}}, false},
		{"HALT", Instruction{HALT, []int{}}, false},
		{"INVALID", Instruction{}, true},
	}

	for _, test := range tests {
		result, err := ParseInstruction(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ParseInstruction(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if result.Opcode != test.expected.Opcode {
			t.Errorf("ParseInstruction(%q) = %v, want %v", test.input, result, test.expected)
		}
		for i := range result.Operands {
			if result.Operands[i] != test.expected.Operands[i] {
				t.Errorf("ParseInstruction(%q) = %v, want %v", test.input, result, test.expected)
			}
		}
	}
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
		{"0", 0, false},
		{"255", 255, false},
		{"256", 0, true},
		{"-1", 0, true},
		{"abc", 0, true},
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
		{"-10", -10, false},
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
