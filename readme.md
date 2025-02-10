# TinyASS - Tiny Assembly Interpreter

TinyASS is a simple emulator for a tiny custom processor and its assembly language. It allows you to write and execute basic assembly instructions.

## Features

- **Hexadecimal Memory Addresses Only:**  
  Memory addresses must be provided in hexadecimal format with a `0x` prefix (e.g., `0x00` to `0xFF`).  
  Decimal values are not accepted.

- **Instructions:**  
  Supports basic arithmetic and bitwise operations, memory store, and branch instructions.  
  Use commands like `LOAD`, `STORE`, `ADD`, etc.

- **REPL Mode and Script Execution:**  
  Run the interpreter in interactive mode or supply a script file.

## Getting Started

### Prerequisites

- Go 1.23.2 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/itsfuad/tinyass.git
    cd tinyass
    ```

2. Build the project:
    ```sh
    go build
    ```

### Usage

- **REPL Mode:**

  ```bash
  go run main.go
  ```

- **Script Mode:**

  ```bash
  go run main.go script.ass
  ```

- **Version:**

  ```bash
  go run main.go --version
  ```

You will enter the REPL interface where you can type assembly instructions.

### Example

```sh
> LOAD R0 10
> LOAD R1 20
> ADD R2 R0 R1
> reg
Registers: R0=10 R1=20 R2=30 R3=0
> HALT
```

## Example Assembly File

```assembly
LOAD R0 10
LOAD R1 0
ADD R2 R0 R1
STORE R2 0x64
DIV R1 R2 R1
PRINT R2      ; Display the value of register R2
PRINT MEM 0x64 ; Display the value at memory address 0x64
```

### Commands

- `LOAD reg val` - Load value into register
- `ADD dest s1 s2` - Add s1 and s2 into dest
- `SUB dest s1 s2` - Subtract s2 from s1 into dest
- `MUL dest s1 s2` - Multiply s1 and s2 into dest
- `DIV dest s1 s2` - Divide s1 by s2 into dest
- `REM dest s1 s2` - Remainder of s1 divided by s2 into dest
- `JMP addr` - Jump to address
- `JZ reg addr` - Jump to address if register is zero
- `JNZ reg addr` - Jump to address if register is not zero
- `HALT` - Stop execution
- `reg` - Show registers
- `exit` - Exit interpreter

## Error Handling

- If an invalid memory address is provided (outside `0x00` to `0xFF`), the interpreter will report an error.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.