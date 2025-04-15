package argsReader

import (
	"flag"
	"strings"
)

// ReadFlags analisa as flags da linha de comando e retorna uma estrutura ConfigArgs
// Se config for fornecida, apenas sobrescreve os valores que foram explicitamente definidos pelas flags
func ReadFlags(initialConfig ConfigArgs) ConfigArgs {
	// Definição dos valores padrão baseados na configuração inicial (se existir)
	compilerPathDefault := ""
	outputNameDefault := ""
	runAfterCompileDefault := false
	customRunCommandDefault := ""
	sourceFileDefault := ""

	// Verifica se uma configuração inicial foi fornecida verificando alguns de seus campos
	configInitialized := initialConfig.CompilerPath != "" // Se pelo menos o caminho do compilador estiver definido

	// Se uma configuração inicial foi fornecida, usa seus valores como padrão
	if configInitialized {
		if initialConfig.CompilerPath != "" {
			compilerPathDefault = initialConfig.CompilerPath
		}
		if initialConfig.OutputName != "" {
			outputNameDefault = initialConfig.OutputName
		}
		runAfterCompileDefault = initialConfig.RunAfterCompile
		if initialConfig.CustomRunCommand != "" {
			customRunCommandDefault = initialConfig.CustomRunCommand
		}
		if initialConfig.SourceFile != "" {
			sourceFileDefault = initialConfig.SourceFile
		}
	}

	// Definição das flags correspondentes aos campos da estrutura ConfigArgs
	compilerPath := flag.String("compiler", compilerPathDefault, "Caminho para o compilador")
	compilerArgsStr := flag.String("args", "", "Argumentos adicionais para compilação (separados por vírgula)")
	outputName := flag.String("output", outputNameDefault, "Nome do arquivo de saída")
	runAfterCompile := flag.Bool("run", runAfterCompileDefault, "Executar após compilar")
	customRunCommand := flag.String("run-cmd", customRunCommandDefault, "Comando personalizado para execução")
	sourceFile := flag.String("source", sourceFileDefault, "Arquivo fonte a ser compilado")

	// Analisa as flags
	flag.Parse()

	// Verifica quais flags foram explicitamente fornecidas pelo usuário
	var compilerArgs []string
	compilerArgsFlagSet := false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "args" && *compilerArgsStr != "" {
			compilerArgsFlagSet = true
			compilerArgs = strings.Split(*compilerArgsStr, ",")
		}
	})

	// Se as flags de argumentos do compilador não foram definidas, mas existem na configuração inicial
	if !compilerArgsFlagSet && configInitialized && len(initialConfig.CompilerArgs) > 0 {
		compilerArgs = initialConfig.CompilerArgs
	}

	// Cria e retorna a instância de ConfigArgs com base nas flags
	config := ConfigArgs{
		CompilerPath:     *compilerPath,
		CompilerArgs:     compilerArgs,
		OutputName:       *outputName,
		RunAfterCompile:  *runAfterCompile,
		CustomRunCommand: *customRunCommand,
		SourceFile:       *sourceFile,
	}

	// Atualiza a variável global Config
	Config = config

	return config
}
