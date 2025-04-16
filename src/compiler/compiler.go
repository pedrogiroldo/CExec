package compiler

import (
	"CExec/src/argsReader"
	"fmt"
	"os"
	"os/exec"
)

func Compile(
	config argsReader.ConfigArgs,
	arquivo string,
	output string,
) bool {
	// Define o comando do compilador e argumentos
	var compilerCmd string
	if config.CompilerPath != "" {
		compilerCmd = config.CompilerPath
	} else {
		fmt.Printf("Compilador não especificado.")
		return false
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
		return false
	}

	return true
}
