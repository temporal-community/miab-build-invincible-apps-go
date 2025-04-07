package durable

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func CountingWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	logger.Info("*** Counting to 10")
	i := 1
	for i <= 10 {
		logger.Info(fmt.Sprintf("%d", i))
		var result int
		err := workflow.ExecuteActivity(ctx, AddOne, i).Get(ctx, &result)
		if err != nil {
			return "", fmt.Errorf("Failed to get location: %s", err)
		}
		i = result
		workflow.Sleep(ctx, 1*time.Second)
	}
	logger.Info("*** Counted to 10")
	return "Counted to 10", nil
}
