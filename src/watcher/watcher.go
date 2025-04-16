package watcher

import (
	"CExec/src/argsReader"
	"CExec/src/compiler"
	"CExec/src/runner"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
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

	// Mensagem informativa sobre como sair
	fmt.Println("\nModo watch iniciado. Pressione Ctrl+C para sair.")

	go func() {
		for {
			select {
			case <-w.Event:
				compiler.Compile(config, file, output)
				runner.Run(config, output)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add("."); err != nil {
		log.Fatalln(err)
	}

	// Rotina para tratar o sinal de interrupção
	go func() {
		<-signalChan
		fmt.Println("\nEncerrando o modo watch...")
		w.Close()
		os.Exit(0)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
