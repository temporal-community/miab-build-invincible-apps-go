package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"miab-build-invincible-apps-go/iplocate"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

var temporalClient client.Client

// Initialize Temporal Client
func initializeTemporal() error {
	var err error
	temporalClient, err = client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	return err
}

// Handle HTMX form submission
func handleSubmit(w http.ResponseWriter, r *http.Request) {
	workflowID := "getAddressFromIP-" + uuid.New().String()
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: iplocate.TaskQueueName,
	}

	name := r.FormValue("name")
	sleepDuration := r.FormValue("sleep_duration")

	var seconds int
	if sleepDuration == "" {
		seconds = 0
	} else {
		// Convert string to integer
		var err error
		seconds, err = strconv.Atoi(sleepDuration)
		if err != nil {
			http.Error(w, "Invalid sleep duration value", http.StatusBadRequest)
			return
		}
	}
	input := iplocate.WorkflowInput{Name: name, Seconds: seconds}

	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, iplocate.GetAddressFromIP, input)
	if err != nil {
		http.Error(w, "An error occurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var result iplocate.WorkflowOutput
	err = we.Get(context.Background(), &result)

	if err != nil {
		http.Error(w, "An error occurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Hello %s. Your IP is %s and your location is %s", input.Name, result.IPAddr, result.Location)
	fmt.Fprintf(w, "<p>%s</p>", message)
}

// Handle HTMX call for demo options
func handleDemoOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%s</p>", iplocate.DemoOptions)
}

// Serve static HTML, CSS, and JS files
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "templates/index.html")
		return
	}
	http.ServeFile(w, r, filepath.Join("templates", r.URL.Path))
}

func main() {
	err := initializeTemporal()
	if err != nil {
		log.Fatalf("Failed to initialize Temporal client: %v", err)
	}

	http.HandleFunc("/greet", handleSubmit)
	http.HandleFunc("/demo-options", handleDemoOptions)
	http.HandleFunc("/", serveStaticFiles)

	port := 8000
	fmt.Printf("Server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
