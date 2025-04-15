package runner

import (
	"CExec/src/argsReader"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
