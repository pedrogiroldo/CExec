package main

import (
	"CExec/src/argsReader"
	"CExec/src/compiler"
	"CExec/src/initializer"
	"CExec/src/runner"
	"CExec/src/watcher"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Verifica se o comando init foi solicitado como subcomando
	if len(os.Args) > 1 && os.Args[1] == "init" {
		initializer.Init()
	}

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
	} else if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") && os.Args[1] != "init" {
		// Se o primeiro argumento não for uma flag ou o comando init, considera como o arquivo fonte
		arquivo = os.Args[1]
	} else {
		// Se nenhum arquivo for especificado, exibe uma mensagem de erro
		fmt.Fprintf(os.Stderr, "No source file specified.\nUsage: %s [file.(c/cpp)] or set the '-source' flag or 'sourceFile' in the configuration file.\n", os.Args[0])
		os.Exit(1)
	}

	// Define o nome do arquivo de saída com base na configuração ou usa o padrão
	output := "output"
	if config.OutputName != "" {
		output = config.OutputName
	} else if os.PathSeparator == '\\' { // Ajusta para Windows
		output = "output.exe"
	}

	if config.WatchChanges {
		watcher.Watch(config, arquivo, output)
	} else {
		if config.RunAfterCompile {
			compiler.Compile(config, arquivo, output)
			runner.Run(config, output)
		} else {
			compiler.Compile(config, arquivo, output)
		}
	}
}
