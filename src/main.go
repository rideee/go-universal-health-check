package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rideee/go-universal-health-check/src/utils"
)

func main() {

	// Measuring execution time.
	startTime := time.Now()

	exec, _ := os.Executable()
	baseDir := filepath.Dir(exec)
	configDir := filepath.Join(baseDir, "config")
	configServersFile := filepath.Join(configDir, "servers.json")

	config := utils.NewConfigParser(configServersFile)

	fmt.Println("\nUniversal Health Check (Go implementation) made by Michał Kątnik (github.com/rideee)")
	fmt.Printf("Config directory location: %s\n", configDir)

	// Main loop.
	for _, srv := range config.ServersConfigData.Servers {
		srvTypeObject := config.ServerTypesMap[srv.ServerType]
		fmt.Printf("\nRunning checks for %s\n", srv.Name)
		fmt.Printf("Type: %s\n", srv.ServerType)
		fmt.Printf("Method: %s, %s:%s logging as %s\n", srvTypeObject.Method, srv.IP, srv.Port, srv.Username)
		fmt.Printf("Command:\n%s\n", srvTypeObject.Command)
	}

	// Elapsed time.
	elapsedTime := time.Since(startTime)
	fmt.Printf("\n\n(Execution time: %s)\n\n", elapsedTime)
}
