package watcher

import (
	"CExec/src/argsReader"
	"CExec/src/compiler"
	"CExec/src/runner"
	"log"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func Watcher() {
	w := watcher.New()
	w.SetMaxEvents(1)

	w.FilterOps(watcher.Write, watcher.Create, watcher.Remove, watcher.Move)

	r := regexp.MustCompile(`.*\\.(c|cpp)$`)

	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case <-w.Event:
				compiler.Compile(argsReader.Config, argsReader.Config.OutputName, argsReader.Config.OutputName)
				runner.Run(argsReader.Config, argsReader.Config.OutputName)
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

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
