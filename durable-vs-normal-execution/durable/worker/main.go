package main

import (
	"log"

	"miab-build-invincible-apps-go/durable-vs-normal-execution/durable"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	// Create the logger with desired level
	customLogger := durable.NewSimpleLogger(1)

	// Create the Temporal client
	c, err := client.Dial(client.Options{
		Logger: customLogger,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create the Temporal worker
	w := worker.New(c, "durable", worker.Options{})

	// Register Workflow and Activities
	w.RegisterWorkflow(durable.CountingWorkflow)
	w.RegisterActivity(durable.AddOne)

	// Start the Worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Temporal worker", err)
	}
}
