package main

import (
	"os"

	"github.com/heroku/rollrus"
	"github.com/jingweno/upterm/cmd/uptermd/internal/command"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	token := os.Getenv("ROLLBAR_ACCESS_TOKEN")
	if token != "" {
		defer rollrus.ReportPanic(token, "uptermd.upterm.dev")
		logger.SetFormatter(&log.TextFormatter{DisableTimestamp: true})
		logger.AddHook(rollrus.NewHook(token, "uptermd.upterm.dev"))
	}

	if err := command.Root(logger).Execute(); err != nil {
		logger.Fatal(err)
	}
}
