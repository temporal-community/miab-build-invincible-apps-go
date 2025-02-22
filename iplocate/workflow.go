package iplocate

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// GetAddressFromIP is the Temporal Workflow that retrieves the IP address and location info.
func GetAddressFromIP(ctx workflow.Context, name string, seconds int) (string, error) {
	// Define the activity options, including the retry policy
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var ipActivities *IPActivities

	var ip string
	err := workflow.ExecuteActivity(ctx, ipActivities.GetIP).Get(ctx, &ip)
	if err != nil {
		return "", fmt.Errorf("Failed to get IP: %s", err)
	}

	workflow.Sleep(ctx, time.Second*time.Duration(seconds))

	var location string
	err = workflow.ExecuteActivity(ctx, ipActivities.GetLocationInfo, ip).Get(ctx, &location)
	if err != nil {
		return "", fmt.Errorf("Failed to get location: %s", err)
	}
	return fmt.Sprintf("Hello, %s. Your IP is %s and your location is %s", name, ip, location), nil
}
