TinyASS - Tiny Assembly Language Interpreter and Emulator

Overview:
------------
TinyASS is an emulator for a custom, minimalistic processor along with its assembly language.
It provides functionality for parsing assembly instructions, executing them on a simulated CPU,
and interacting via a command-line REPL or by running script files. The project is structured
into several packages, each with distinct responsibilities:

1. commands:
  - Contains definitions for opcodes and the Instruction type that represents a single assembly
    instruction. 
  - Implements parsing functions (e.g., ParseInstruction, ParseRegister, ParseMemory, ParseValue)
    to convert string representations of instructions into structured data.
  - Provides error handling with detailed messages for invalid registers, memory addresses, and
    values to ensure robust input validation.
  - Supports a comprehensive set of operations including arithmetic (ADD, SUB, MUL, DIV, REM),
    bitwise (AND, OR, XOR, NOT), shift instructions (SHL, SHR), comparisons (GT, LT, GTE, LTE, EQ, NEQ),
    jumps (JMP, JZ, JNZ), memory operations (LOAD, STORE), and output (PRINT, HALT).

2. runtime:
  - Implements the CPU simulation which contains registers, memory, the program counter, and the
    loaded instructions.
  - Provides methods (e.g., LoadProgram, Execute) for loading a parsed program into CPU memory and
    executing each instruction sequentially.
  - Handles execution flow control, including jump instructions and error detection (e.g., division
    by zero), making sure the CPU state is accurately maintained.
  - Offers a REPL (Read-Eval-Print Loop) mode so users can manually enter and execute assembly
    commands interactively.
  - Includes auxiliary functions to display the CPU’s registers and memory state in a user-friendly
    manner with color-coded output.

3. utils:
  - Supplies helper functionality for better user experience such as colored terminal output using
    ANSI escape codes.
  - Contains utility functions like ClearScreen to support cross-platform console manipulations.
  - Enhances overall readability and maintenance by centralizing common behaviors like formatted
    print routines and error messaging.

4. main:
  - Acts as the entry point for the TinyASS application.
  - Parses command-line arguments to determine whether to run a provided assembly script (via RunFile)
    or to launch the interactive REPL environment.
  - Provides version information to users through a command-line flag.

5. CI/CD Workflows:
  - The repository includes GitHub Actions workflows for continuous integration (CI) that compile,
    lint, format check, and test the Go code on each push or pull request.
  - A separate release workflow ensures automated building of binaries for multiple platforms,
    version tagging, and GitHub Releases integration for streamlined deployment.

Testing:
------------
A comprehensive suite of unit tests validates the functionality of instruction parsing and CPU
execution (e.g., handling arithmetic operations, error conditions like division by zero, and output
via PRINT). This helps maintain code quality and ensures that updates don’t introduce regressions.

Usage:
------------
Users can interact with TinyASS either by running assembly scripts from files or by entering commands
directly into the REPL. Commands support essential operations to emulate the behavior of simple
assembly instructions, making the tool ideal for educational purposes or as a basis for further
extension and experimentation with low-level programming concepts.

Overall, TinyASS is designed to be modular, extensible, and user-friendly, providing clear error messages,
robust parsing, and a straightforward emulation environment for custom assembly language programming.

## Example Usage

Download the release binary for your platform from the [Releases](github.com/itsfuad/tinyass/releases) page.

To build the project:
```bash
go build -o tinyass main.go
```

To run an assembly script:
```bash
go run main.go path/to/script.ass
```

To launch the interactive REPL:
```bash
go run main.go
```

Display version information:
```bash
go run main.go --version
```

## Contribution Guidelines

Contributions are welcome! Please consider the following:
- Fork the repository and create a new branch for your feature or bugfix.
- Follow the existing coding style and include tests where applicable.
- Ensure your changes pass the CI pipeline.
- Update documentation and example usage if needed.
- Run `go fmt ./...` before pushing.
- Open a pull request with a clear description of your changes and reference any related issues.
- For major changes, please open an issue first to discuss what you would like to change.

All contributions must adhere to the project’s [License](LICENSE).