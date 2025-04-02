package main

import (
	"context"
	"fmt"
	"log"

	"miab-build-invincible-apps-go/durable-vs-volatile/durable"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowID := fmt.Sprintf("getAddressFromIP-" + uuid.New().String())

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "durable",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, durable.CountingWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
