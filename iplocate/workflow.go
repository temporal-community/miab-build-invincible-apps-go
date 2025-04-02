package iplocate

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// GetAddressFromIP is the Temporal Workflow that retrieves the IP address and location info.
func GetAddressFromIP(ctx workflow.Context, input WorkflowInput) (WorkflowOutput, error) {
	// Define the activity options, including the retry policy
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var ip string
	err := workflow.ExecuteActivity(ctx, GetIP).Get(ctx, &ip)
	if err != nil {
		return WorkflowOutput{}, fmt.Errorf("Failed to get IP: %s", err)
	}

	if input.Seconds > 0 {
		workflow.Sleep(ctx, time.Second*time.Duration(input.Seconds))
	}

	var location string
	err = workflow.ExecuteActivity(ctx, GetLocationInfo, ip).Get(ctx, &location)
	if err != nil {
		return WorkflowOutput{}, fmt.Errorf("Failed to get location: %s", err)
	}
	return WorkflowOutput{IPAddr: ip, Location: location}, nil
	//return fmt.Sprintf("Hello, %s. Your IP is %s and your location is %s", name, ip, location), nil
}
