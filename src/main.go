package main

import (
	"CExec/src/fileReader"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	config := fileReader.ReadFile()

	// Determina o arquivo fonte a ser compilado
	var arquivo string
	if len(os.Args) >= 2 {
		// Se um arquivo for fornecido via linha de comando, ele tem precedência
		arquivo = os.Args[1]
	} else if config.SourceFile != "" {
		// Caso contrário, usa o arquivo especificado na configuração
		arquivo = config.SourceFile
	} else {
		// Se nenhum arquivo for especificado, exibe uma mensagem de erro
		fmt.Fprintf(os.Stderr, "Nenhum arquivo fonte especificado.\nUso: %s [arquivo.(c/cpp)] ou defina 'sourceFile' no arquivo de configuração.\n", os.Args[0])
		os.Exit(1)
	}

	// Define o nome do arquivo de saída com base na configuração ou usa o padrão
	output := "output"
	if config.OutputName != "" {
		output = config.OutputName
	} else if os.PathSeparator == '\\' { // Ajusta para Windows
		output = "output.exe"
	}

	// Define o comando do compilador e argumentos
	var compilerCmd string
	if config.CompilerPath != "" {
		compilerCmd = config.CompilerPath
	} else {
		fmt.Printf("Compilador não especificado.")
		os.Exit(1)

	}

	// Prepara os argumentos de compilação
	compilerArgs := []string{"-o", output, arquivo}
	if len(config.CompilerArgs) > 0 {
		compilerArgs = append([]string{"-o", output}, config.CompilerArgs...)
		compilerArgs = append(compilerArgs, arquivo)
	}

	// Executa a compilação
	compilacao := exec.Command(compilerCmd, compilerArgs...)
	compilacao.Stderr = os.Stderr
	if err := compilacao.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na compilação: %v\n", err)
		os.Exit(1)
	}

	// Executa o programa compilado se configurado para isso
	if config.RunAfterCompile {
		execPath := "." + string(os.PathSeparator) + output

		if config.CustomRunCommand != "" {
			// Separa o comando e os argumentos
			args := strings.Fields(config.CustomRunCommand)
			execucao := exec.Command(execPath, args...)
			execucao.Stdout = os.Stdout
			execucao.Stderr = os.Stderr
			execucao.Stdin = os.Stdin
			if err := execucao.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Erro na execução: %v\n", err)
				os.Exit(1)
			}
		} else {
			// Sem argumentos adicionais
			execucao := exec.Command(execPath)
			execucao.Stdout = os.Stdout
			execucao.Stderr = os.Stderr
			execucao.Stdin = os.Stdin
			if err := execucao.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Erro na execução: %v\n", err)
				os.Exit(1)
			}
		}
	}
}
