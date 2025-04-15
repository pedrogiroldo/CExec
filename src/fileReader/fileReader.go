package fileReader

import (
	"encoding/json"
	"fmt"
	"os"
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

func ReadFile() ConfigFile {
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

	return config

}
