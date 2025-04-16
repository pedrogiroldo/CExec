package main

import (
	"CExec/src/argsReader"
	"CExec/src/compiler"
	"CExec/src/runner"
	"CExec/src/watcher"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Primeiro, tenta ler o arquivo de configuração (se existir)
	var config argsReader.ConfigArgs
	fileExists := argsReader.FileExists()

	if fileExists {
		config = argsReader.ReadFile()
	}

	// 2. Em seguida, lê as flags da linha de comando (que sobrescrevem as configurações do arquivo)
	config = argsReader.ReadFlags(config)

	// Determina o arquivo fonte a ser compilado
	var arquivo string
	if config.SourceFile != "" {
		// Usa o arquivo especificado via flag ou configuração
		arquivo = config.SourceFile
	} else if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		// Se o primeiro argumento não for uma flag, considera como o arquivo fonte
		arquivo = os.Args[1]
	} else {
		// Se nenhum arquivo for especificado, exibe uma mensagem de erro
		fmt.Fprintf(os.Stderr, "Nenhum arquivo fonte especificado.\nUso: %s [arquivo.(c/cpp)] ou defina a flag '-source' ou 'sourceFile' no arquivo de configuração.\n", os.Args[0])
		os.Exit(1)
	}

	// Define o nome do arquivo de saída com base na configuração ou usa o padrão
	output := "output"
	if config.OutputName != "" {
		output = config.OutputName
	} else if os.PathSeparator == '\\' { // Ajusta para Windows
		output = "output.exe"
	}

	compiler.Compile(config, arquivo, output)

	if config.WatchChanges {
		watcher.Watch(config, arquivo, output)
	} else {
		if config.RunAfterCompile {
			runner.Run(config, output)
		}

	}

}
