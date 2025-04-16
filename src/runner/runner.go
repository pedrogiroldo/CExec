package runner

import (
	"CExec/src/argsReader"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run executa o programa compilado de forma síncrona (bloqueante)
func Run(config argsReader.ConfigArgs, output string) {
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

// StartAsync inicia o programa compilado de forma assíncrona (não bloqueante)
// Retorna o comando em execução que pode ser gerenciado pelo chamador
func StartAsync(config argsReader.ConfigArgs, output string) (*exec.Cmd, error) {
	execPath := "." + string(os.PathSeparator) + output
	var cmd *exec.Cmd

	if config.CustomRunCommand != "" {
		// Separa o comando e os argumentos
		args := strings.Fields(config.CustomRunCommand)
		cmd = exec.Command(execPath, args...)
	} else {
		// Sem argumentos adicionais
		cmd = exec.Command(execPath)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
