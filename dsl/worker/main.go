package main

import (
	"go.temporal.io/temporal/client"
	"go.temporal.io/temporal/worker"
	"go.uber.org/zap"

	"github.com/temporalio/temporal-go-samples/dsl"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal("Unable to create client", zap.Error(err))
	}
	defer c.Close()

	w := worker.New(c, "dsl", worker.Options{})
	defer w.Stop()

	w.RegisterWorkflow(dsl.SimpleDSLWorkflow)
	w.RegisterActivity(&dsl.SampleActivities{})

	err = w.Run()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}
}
