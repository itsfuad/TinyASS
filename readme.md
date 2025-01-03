# TinyASS - Tiny Assembly Interpreter

TinyASS is a simple emulator for a tiny custom processor and its assembly language. It allows you to write and execute basic assembly instructions.

## Features

- Supports basic arithmetic operations: ADD, SUB, MUL, DIV, REM
- Memory operations: LOAD, STORE
- Control flow instructions: JMP, JZ, JNZ, HALT
- Simple REPL interface for interactive use

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

Run the interpreter:
```sh
./tinyass
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

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.