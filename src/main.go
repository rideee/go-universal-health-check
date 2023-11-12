package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rideee/go-universal-health-check/src/utils"
)

func main() {

	exec, _ := os.Executable()
	baseDir := filepath.Dir(exec)
	configDir := filepath.Join(baseDir, "config")
	configServersFile := filepath.Join(configDir, "servers.json")

	config := utils.NewConfigParser(configServersFile)

	fmt.Println("\nUniversal Health Check (Go implementation) made by Michał Kątnik (github.com/rideee)")
	fmt.Printf("Config directory location: %s\n\n", configDir)

	// Main loop.
	for _, srv := range config.ServersConfigData.Servers {
		fmt.Printf("Running checks for %s\n", srv.Name)
	}

}
