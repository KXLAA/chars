package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"github.com/KXLAA/chars/pkg/leveledlog"
	"github.com/KXLAA/chars/pkg/version"
)

type application struct {
	config config
	logger *leveledlog.Logger
	wg     sync.WaitGroup
}

func main() {
	logger := leveledlog.NewLogger(os.Stdout, leveledlog.LevelAll, true)
	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatal(err, trace)
	}

}

type config struct {
	baseURL  string
	httpPort int
}

func run(logger *leveledlog.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 4444, "port to listen on for HTTP requests")
	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	return app.serveHTTP()
}
