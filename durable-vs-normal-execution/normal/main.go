package main

import (
	"log"
	"time"
)

func init() {
	// Configure logging to only print the message
	log.SetFlags(0)
}

func main() {
	for i := 1; i <= 10; i++ {
		log.Printf("%d", i)
		time.Sleep(1 * time.Second)
	}
}
