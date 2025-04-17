package initializer

import (
	"CExec/src/argsReader"
	"fmt"
	"os"
	"strings"
)

func Init() {

	var compilerPath, outputName, runAfterCompile, sourceFile, watchChanges string

	fmt.Println("CExec Configuration:")
	fmt.Println("Enter the compiler path (e.g.: g++):")
	fmt.Scan(&compilerPath)

	fmt.Println("Enter the output file name (without extension):")
	fmt.Scan(&outputName)

	fmt.Println("Run the program after compilation? (y/n):")
	fmt.Scan(&runAfterCompile)

	fmt.Println("Enter the source file name (e.g.: main.cpp):")
	fmt.Scan(&sourceFile)

	fmt.Println("Watch file changes? (y/n):")
	fmt.Scan(&watchChanges)

	// Configuração padrão para o comando init
	defaultConfig := argsReader.ConfigArgs{
		CompilerPath:     compilerPath,
		CompilerArgs:     []string{},
		OutputName:       outputName,
		RunAfterCompile:  runAfterCompile == "s" || runAfterCompile == "sim" || runAfterCompile == "y" || runAfterCompile == "yes",
		CustomRunCommand: "",
		SourceFile:       sourceFile,
		WatchChanges:     watchChanges == "s" || watchChanges == "sim" || watchChanges == "y" || watchChanges == "yes",
	}

	// Verifica se já existe um arquivo de configuração
	fileExists := argsReader.FileExists()
	if fileExists {
		fmt.Println("Configuration file already exists. Do you want to overwrite it? (y/n):")
		var resposta string
		fmt.Scanln(&resposta)
		resposta = strings.ToLower(resposta)

		if resposta != "s" && resposta != "sim" && resposta != "y" && resposta != "yes" {
			fmt.Println("Operation cancelled.")
			os.Exit(0)
		}
	}

	// Salva o arquivo de configuração
	err := argsReader.SaveConfigFile(defaultConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating configuration file: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
