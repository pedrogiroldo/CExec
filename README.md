# CExec

A command-line tool to compile and run C/C++ files with a single command.

[🇧🇷 Portuguese Version (Versão em Português)](README.pt-br.md)

## About

CExec is a lightweight GO tool crafted to streamline the compilation and execution of C/C++ programs. With just one terminal command, it automates the entire process, making it effortless for developers to test and run their code. Designed for efficiency, CExec simplifies workflows by combining both steps into a single, intuitive action.

## Traditional Approach vs CExec

### Traditional Approach

The traditional way requires multiple commands to compile and then run C/C++ programs:

![Traditional C/C++ Compilation and Execution](docs/assets/ex1.gif)

### CExec Approach

With CExec, the entire process is simplified into a single command:

![CExec Compilation and Execution](docs/assets/ex2.gif)

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

### Command-line options

You can also use command-line flags to customize CExec's behavior:

```bash
CExec -compiler=/path/to/g++ -args="-Wall,-std=c++17" -output=my_program -run=true -source=main.cpp -watch=true
```

| Flag        | Description                                          | Default                            |
| ----------- | ---------------------------------------------------- | ---------------------------------- |
| `-compiler` | Path to the compiler                                 | From config file or required       |
| `-args`     | Compiler arguments (comma-separated)                 | From config file or none           |
| `-output`   | Name of the output executable                        | "output" or "output.exe" (Windows) |
| `-run`      | Whether to run the program after compilation         | From config file or false          |
| `-run-cmd`  | Custom arguments to pass to the program when running | From config file or none           |
| `-source`   | Source file to compile                               | From config file or required       |
| `-watch`    | Enable file watching mode to recompile on changes    | From config file or false          |

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
  "sourceFile": "main.cpp",
  "watchChanges": false
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
| `watchChanges`     | Enable file watching mode to recompile on changes        | No (default: false)                               |

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

### Watch mode example

```bash
$ CExec -watch=true my_program.cpp
```

This will start CExec in watch mode. It will compile the file, and then continue monitoring it for changes. Whenever the file is modified and saved, CExec will automatically recompile it.

## Compatibility notes

The executable automatically detects the operating system and adjusts the output file name:

- On Linux/Unix systems: `output`
- On Windows systems: `output.exe`

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests.
