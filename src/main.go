package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Verifica se foi fornecido o argumento correto
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s <arquivo.cpp>\n", os.Args[0])
		os.Exit(1)
	}

	// Armazena o nome do arquivo fornecido
	arquivo := os.Args[1]
	output := "output"
	if os.PathSeparator == '\\' { // Ajusta para Windows
		output = "output.exe"
	}

	// Executa a compilação
	compilacao := exec.Command("g++", "-o", output, arquivo)
	compilacao.Stderr = os.Stderr
	if err := compilacao.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na compilação: %v\n", err)
		os.Exit(1)
	}

	// Executa o programa compilado
	execucao := exec.Command("." + string(os.PathSeparator) + output)
	execucao.Stdout = os.Stdout
	execucao.Stderr = os.Stderr
	if err := execucao.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na execução: %v\n", err)
		os.Exit(1)
	}
	
}