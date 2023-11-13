package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rideee/go-universal-health-check/src/models"
	"github.com/rideee/go-universal-health-check/src/utils"
)

func main() {

	// Measuring execution time.
	startTime := time.Now()

	// Set Paths.
	exec, _ := os.Executable()
	baseDir := filepath.Dir(exec)
	configDir := filepath.Join(baseDir, "config")
	configServersFile := filepath.Join(configDir, "servers.json")

	config := utils.NewConfigParser(configServersFile)

	fmt.Println("\nUniversal Health Check (Go implementation) made by Michał Kątnik (github.com/rideee)")
	fmt.Printf("Config directory location: %s\n", configDir)

	// Initialize channels slice.
	var channels []chan string

	// Goroutine main loop.
	for _, srv := range config.ServersConfigData.Servers {

		// Make a new channel for every iterration of the loop and apped it to channels slice.
		channel := make(chan string)
		channels = append(channels, channel)

		// Run goroutine annonymous function.
		go func(channel chan string, srv models.Server) {

			srvTypeObject := config.ServerTypesMap[srv.ServerType]
			output := fmt.Sprintf("\nRunning checks for %s\n", srv.Name)
			output += fmt.Sprintf("Type: %s\n", srv.ServerType)
			output += fmt.Sprintf("Method: %s, %s:%s logging as %s\n", srvTypeObject.Method, srv.IP, srv.Port, srv.Username)
			output += fmt.Sprintf("Command:\n%s\n", srvTypeObject.Command)

			channel <- output

		}(channel, srv)
	}

	// Collect data from all channels in correct order (FIFO).
	for _, ch := range channels {
		fmt.Println(<-ch)
	}

	// Elapsed time.
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\n(Execution time: %s)\n\n", elapsedTime)
}
