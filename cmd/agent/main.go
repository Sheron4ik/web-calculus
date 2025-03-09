package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Sheron4ik/web-calculus/internal/agent"
	"github.com/Sheron4ik/web-calculus/internal/app"
)

func main() {
	agentApp := app.New()

	computingPower, _ := strconv.Atoi(agentApp.Config.ComputingPower)

	for i := range computingPower {
		go agent.Worker(i, agentApp.Config.Port)
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)
	<-gracefulShutdown
}
