package argsReader

type ConfigArgs struct {
	CompilerPath     string   `json:"compilerPath"`     // Caminho para o compilador (g++)
	CompilerArgs     []string `json:"compilerArgs"`     // Argumentos adicionais para compilação
	OutputName       string   `json:"outputName"`       // Nome do arquivo de saída (padrão: output/output.exe)
	RunAfterCompile  bool     `json:"runAfterCompile"`  // Executar após compilar
	CustomRunCommand string   `json:"customRunCommand"` // Comando personalizado para execução
	SourceFile       string   `json:"sourceFile"`       // Arquivo fonte a ser compilado
	WatchChanges     bool     `json:"watchChanges"`     // Habilitar monitoramento de alterações no arquivo fonte
}
