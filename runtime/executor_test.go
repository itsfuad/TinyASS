package runtime

import (
	"testing"
	"tinyass/commands"
)

func TestCPUExecute(t *testing.T) {
	tests := []struct {
		name        string
		instruction commands.Instruction
		initialCPU  CPU
		expectedCPU CPU
		expectError bool
	}{
		{
			name:        "LOAD",
			instruction: commands.Instruction{Opcode: commands.LOAD, Operands: []int{0, 10}},
			initialCPU:  CPU{},
			expectedCPU: CPU{registers: [4]int{10, 0, 0, 0}},
		},
		{
			name:        "ADD",
			instruction: commands.Instruction{Opcode: commands.ADD, Operands: []int{2, 0, 1}},
			initialCPU:  CPU{registers: [4]int{10, 20, 0, 0}},
			expectedCPU: CPU{registers: [4]int{10, 20, 30, 0}},
		},
		{
			name:        "DIV by zero",
			instruction: commands.Instruction{Opcode: commands.DIV, Operands: []int{2, 0, 1}},
			initialCPU:  CPU{registers: [4]int{10, 0, 0, 0}},
			expectedCPU: CPU{registers: [4]int{10, 0, 0, 0}},
			expectError: true,
		},
		{
			name:        "PRINT register",
			instruction: commands.Instruction{Opcode: commands.PRINT, Operands: []int{-1, 0}},
			initialCPU:  CPU{registers: [4]int{10, 0, 0, 0}},
			expectedCPU: CPU{registers: [4]int{10, 0, 0, 0}},
		},
		{
			name:        "PRINT memory",
			instruction: commands.Instruction{Opcode: commands.PRINT, Operands: []int{100}},
			initialCPU:  CPU{memory: [commands.MEMORY_SIZE]int{100: 42}},
			expectedCPU: CPU{memory: [commands.MEMORY_SIZE]int{100: 42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := tt.initialCPU
			cpu.Execute(tt.instruction)
		})
	}
}

func TestCPULoadProgram(t *testing.T) {
	cpu := NewCPU()
	program := []commands.Instruction{
		{Opcode: commands.LOAD, Operands: []int{0, 10}},
		{Opcode: commands.ADD, Operands: []int{2, 0, 1}},
	}
	cpu.LoadProgram(program)
	if len(cpu.program) != len(program) {
		t.Errorf("CPU.LoadProgram() = %v, want %v", len(cpu.program), len(program))
	}
	for i, inst := range cpu.program {
		if inst.Opcode != program[i].Opcode {
			t.Errorf("CPU.LoadProgram() = %v, want %v", inst, program[i])
		}
	}
}
