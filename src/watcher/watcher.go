package watcher

import (
	"CExec/src/argsReader"
	"CExec/src/compiler"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/radovskyb/watcher"
)

func Watch(config argsReader.ConfigArgs, file string, output string) {
	w := watcher.New()
	w.SetMaxEvents(1)

	w.FilterOps(watcher.Write, watcher.Create, watcher.Remove, watcher.Move)

	r := regexp.MustCompile(`.*\.(c|cpp)$`)

	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	// Canal para capturar sinais de interrupção (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Variável para armazenar o processo em execução
	var currentCmd *exec.Cmd

	// Função para encerrar o processo atual se existir
	killCurrentProcess := func() {
		if currentCmd != nil && currentCmd.Process != nil {
			// Tenta encerrar o processo de forma limpa
			_ = currentCmd.Process.Signal(syscall.SIGTERM)

			// Aguarda um curto tempo para encerramento limpo
			time.Sleep(100 * time.Millisecond)

			// Caso ainda esteja rodando, força o encerramento
			_ = currentCmd.Process.Kill()
			currentCmd = nil
		}
	}

	// Mensagem informativa sobre como sair
	fmt.Println("\nModo watch iniciado. Pressione Ctrl+C para sair.")

	go func() {
		for {
			select {
			case <-w.Event:
				// Encerra o processo atual antes de recompilar
				killCurrentProcess()

				// Compila o arquivo
				if compiler.Compile(config, file, output) {
					// Executa o programa sem bloquear o watcher
					execPath := "." + string(os.PathSeparator) + output

					if config.CustomRunCommand != "" {
						args := strings.Fields(config.CustomRunCommand)
						currentCmd = exec.Command(execPath, args...)
					} else {
						currentCmd = exec.Command(execPath)
					}

					currentCmd.Stdout = os.Stdout
					currentCmd.Stderr = os.Stderr
					currentCmd.Stdin = os.Stdin

					if err := currentCmd.Start(); err != nil {
						fmt.Fprintf(os.Stderr, "Erro ao iniciar a execução: %v\n", err)
					} else {
						fmt.Println("Programa iniciado. Monitorando mudanças...")

						// Monitora o término do processo sem bloquear
						go func(cmd *exec.Cmd) {
							_ = cmd.Wait()
						}(currentCmd)
					}
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive("."); err != nil {
		log.Fatalln(err)
	}

	// Rotina para tratar o sinal de interrupção
	go func() {
		<-signalChan
		fmt.Println("\nEncerrando o modo watch...")
		killCurrentProcess()
		w.Close()
		os.Exit(0)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
