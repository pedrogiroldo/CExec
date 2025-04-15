package argsReader

type ConfigArgs struct {
	CompilerPath     string   `json:"compilerPath"`               // Caminho para o compilador (g++)
	CompilerArgs     []string `json:"compilerArgs,omitempty"`     // Argumentos adicionais para compilação
	OutputName       string   `json:"outputName,omitempty"`       // Nome do arquivo de saída (padrão: output/output.exe)
	RunAfterCompile  bool     `json:"runAfterCompile"`            // Executar após compilar
	CustomRunCommand string   `json:"customRunCommand,omitempty"` // Comando personalizado para execução
	SourceFile       string   `json:"sourceFile,omitempty"`       // Arquivo fonte a ser compilado
}
