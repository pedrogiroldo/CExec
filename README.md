# CExec

A command-line tool to compile and run C/C++ files with a single command.

[🇧🇷 Portuguese Version (Versão em Português)](README.pt-br.md)

## About

CExec is a lightweight GO tool crafted to streamline the compilation and execution of C/C++ programs. With just one terminal command, it automates the entire process, making it effortless for developers to test and run their code. Designed for efficiency, CExec simplifies workflows by combining both steps into a single, intuitive action.

## Prerequisites

- Go (to build from source)
- A C/C++ compiler (such as g++)

## Installation

### From source

1. Clone this repository
2. Compile the program:

```bash
go build -o build/CExec src/main.go
```

3. Add the compiled executable to your PATH for global usage (optional)

## Usage

### Basic usage

```bash
CExec file.(c/cpp)
```

The program will:

1. Compile the specified C/C++ file using the configured compiler
2. Run the resulting program (if configured)
3. Display the program output

### Configuration file

CExec can be configured through a JSON file called `CExecConfig.json`. This file should be in the same directory from where CExec is executed.

Example of `CExecConfig.json`:

```json
{
  "compilerPath": "/usr/bin/g++",
  "compilerArgs": ["-Wall", "-std=c++17"],
  "outputName": "my_program",
  "runAfterCompile": true,
  "customRunCommand": "arg1 arg2",
  "sourceFile": "main.cpp"
}
```

#### Configuration options:

| Option             | Description                                              | Required                                          |
| ------------------ | -------------------------------------------------------- | ------------------------------------------------- |
| `compilerPath`     | Path to the compiler (e.g., g++)                         | Yes                                               |
| `compilerArgs`     | List of arguments for the compiler                       | No                                                |
| `outputName`       | Name of the generated executable file                    | No (default: "output" or "output.exe" on Windows) |
| `runAfterCompile`  | Whether the program should be executed after compilation | No (default: false)                               |
| `customRunCommand` | Arguments to pass to the program during execution        | No                                                |
| `sourceFile`       | Default source file to be compiled                       | No (can be overridden via command line)           |

**Note:** When a file is specified via the command line, it takes precedence over the file defined in the configuration.

## Example

```bash
$ CExec my_program.cpp
```

If the program `my_program.cpp` contains:

```cpp
#include <iostream>

int main() {
    std::cout << "Hello, world!" << std::endl;
    return 0;
}
```

The output will be:

```
Hello, world!
```

## Compatibility notes

The executable automatically detects the operating system and adjusts the output file name:

- On Linux/Unix systems: `output`
- On Windows systems: `output.exe`

## License

[Add license information]

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests.
