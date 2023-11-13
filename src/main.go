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

	// Measure execution time.
	startTime := time.Now()

	// Set Paths.
	exec, _ := os.Executable()
	baseDir := filepath.Dir(exec)
	configDir := filepath.Join(baseDir, "config")
	configServersFile := filepath.Join(configDir, "servers.json")

	// Initialize config parser.
	config := utils.NewConfigParser(configServersFile)

	// Set colors.
	// TODO: get the colors from general.json config file.
	const COLOR_RESET string = "\u001B[0m"
	const COLOR_PRIMARY string = "\u001B[35m"
	const COLOR_SECONDARY string = "\u001B[36m"
	const COLOR_TERTIARY string = "\u001B[34m"
	const COLOR_WARNING string = "\u001B[33m"
	const COLOR_ERROR string = "\u001B[31m"

	fmt.Print(COLOR_SECONDARY)
	fmt.Println("\nUniversal Health Check (Go implementation) made by Michał Kątnik (github.com/rideee/go-universal-health-check)")
	fmt.Print(COLOR_TERTIARY)
	fmt.Printf("Config directory location: %s\n", configDir)
	fmt.Print(COLOR_RESET)

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
			output := COLOR_PRIMARY
			output += fmt.Sprintf("\n%v | %s ", time.Now().Format("2006-01-02 15:04:05"), COLOR_SECONDARY+srv.Name+COLOR_PRIMARY)
			output += fmt.Sprintf("| Type: %s; ", srv.ServerType)
			output += fmt.Sprintf("Method: %s, %s:%s logging as %s\n", srvTypeObject.Method, srv.IP, srv.Port, srv.Username)
			output += COLOR_RESET
			output += fmt.Sprintf("Command:\n%s\n", srvTypeObject.Command)

			// Send data to the channel.
			channel <- output

		}(channel, srv)
	}

	// Collect data from all channels in correct order (FIFO) and print them.
	for _, ch := range channels {
		fmt.Println(<-ch)
	}

	// Elapsed time.
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\n%s(Execution time: %s)%s\n\n", COLOR_TERTIARY, elapsedTime, COLOR_RESET)
}
