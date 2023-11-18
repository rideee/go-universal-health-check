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

	// Set paths.
	exec, _ := os.Executable()
	baseDir := filepath.Dir(exec)
	configDir := filepath.Join(baseDir, "config")

	// Initialize config parser.
	config := utils.NewConfigParser(configDir)

	// Initialize logger.
	logger := utils.NewLogger(config.GeneralConfigData.DebugMode)

	logger.Debugf("Startup time: %v\n", time.Now().Format(time.UnixDate))
	logger.Debugf("PATHS:\nEXEC: %v\nBASE_DIR: %v\nCONFIG_DIR: %v\n", exec, baseDir, configDir)

	// Set colors.
	logger.Debug("Assigning color variables from a JSON file.")
	logger.Debugf("ColorfulOutput set to %v\n", config.GeneralConfigData.ColorfulOutput)
	const COLOR_RESET string = "\u001B[0m"
	var COLOR_PRIMARY string
	var COLOR_SECONDARY string
	var COLOR_TERTIARY string
	if config.GeneralConfigData.ColorfulOutput {
		COLOR_PRIMARY = config.GeneralConfigData.ColorPrimary
		COLOR_SECONDARY = config.GeneralConfigData.ColorSecondary
		COLOR_TERTIARY = config.GeneralConfigData.ColorTertiary
	}

	// Welcome message.
	logger.Debug("Printing welcome message.")
	fmt.Print(COLOR_SECONDARY)
	fmt.Println("\nUniversal Health Check (Go implementation) made by Michał Kątnik (github.com/rideee/go-universal-health-check)")
	fmt.Print(COLOR_TERTIARY)
	fmt.Printf("Config directory location: %s\n", configDir)
	fmt.Print(COLOR_RESET)

	// Initialize channels slice.
	logger.Debug("Initializing channels slice.")
	var channels []chan string

	// Goroutine main loop.
	logger.Debug("Starting the main goroutines loop.")
	for i, srv := range config.ServersConfigData.Servers {

		logger.Debugf("Goroutine loop - iteration no. %d (for %s).\n", i+1, srv)

		// Make a new channel for every iterration of the loop and apped it to channels slice.
		logger.Debug("Making a new channel.")
		channel := make(chan string)
		logger.Debug("Appending a channel to a channel slice.")
		channels = append(channels, channel)

		// Run goroutine annonymous function.
		logger.Debug("Running goroutine func.")
		go func(channel chan string, srv models.Server) {

			srvTypeObject := config.ServerTypesMap[srv.ServerType]
			output := COLOR_PRIMARY
			output += fmt.Sprintf("\n%v | %s ", time.Now().Format("2006-01-02 15:04:05"), COLOR_SECONDARY+srv.Name+COLOR_PRIMARY)
			output += fmt.Sprintf("| Type: %s; ", srv.ServerType)
			output += fmt.Sprintf("Method: %s, %s:%s logging as %s\n", srvTypeObject.Method, srv.IP, srv.Port, srv.Username)
			output += COLOR_RESET

			// TODO: Implement real action.
			output += fmt.Sprintf("Command:\n%s\n", srvTypeObject.Command)

			// Send data to the channel.
			logger.Debug("Sending data from goroutine func to the channel.")
			channel <- output

		}(channel, srv)
	}

	// Collect data from all channels in correct order (FIFO) and print them.
	logger.Debug("Starting a channel loop.")
	for _, ch := range channels {
		logger.Debug("Printing data from the channel.")
		fmt.Println(<-ch)
	}

	// Elapsed time.
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\n%s(Execution time: %s)%s\n\n", COLOR_TERTIARY, elapsedTime, COLOR_RESET)
}
