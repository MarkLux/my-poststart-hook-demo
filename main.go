package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		status := os.Getenv("CONTAINER_STATUS")
		fmt.Printf("CHECKING CONTAINER_STATUS: %s\n", status)
		if status == "Ready" {
			fmt.Printf("CONTAINER_STATUS IS READY")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
