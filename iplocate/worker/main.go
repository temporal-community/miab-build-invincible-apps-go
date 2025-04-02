package main

import (
	"log"

	"miab-build-invincible-apps-go/iplocate"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Create the Temporal worker
	w := worker.New(c, iplocate.TaskQueueName, worker.Options{})

	// Register Workflow and Activities
	w.RegisterWorkflow(iplocate.GetAddressFromIP)
	w.RegisterActivity(iplocate.GetIP)
	w.RegisterActivity(iplocate.GetLocationInfo)

	// Start the Worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Temporal worker", err)
	}
}
