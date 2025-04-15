package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ConfigFile representa a estrutura do arquivo de configuração JSON
type ConfigFile struct {
	CompilerPath     string   `json:"compilerPath"`               // Caminho para o compilador (g++)
	CompilerArgs     []string `json:"compilerArgs,omitempty"`     // Argumentos adicionais para compilação
	OutputName       string   `json:"outputName,omitempty"`       // Nome do arquivo de saída (padrão: output/output.exe)
	RunAfterCompile  bool     `json:"runAfterCompile"`            // Executar após compilar
	CustomRunCommand string   `json:"customRunCommand,omitempty"` // Comando personalizado para execução
	SourceFile       string   `json:"sourceFile,omitempty"`       // Arquivo fonte a ser compilado
}

func main() {
	// Constant declaration
	const configFile = "CExecConfig.json"

	// Variable declaration with error handling
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening config file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decodifica o JSON para a estrutura ConfigFile
	var config ConfigFile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao decodificar o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

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
