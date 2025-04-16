package argsReader

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFilePath = "CExecConfig.json"

// FileExists verifica se o arquivo de configuração existe
func FileExists() bool {
	_, err := os.Stat(ConfigFilePath)
	return err == nil
}

func ReadFile() ConfigArgs {
	// Variable declaration with error handling
	file, err := os.Open(ConfigFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening config file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decodifica o JSON para a estrutura ConfigArgs
	var config ConfigArgs
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao decodificar o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

	Config = config

	return config
}
