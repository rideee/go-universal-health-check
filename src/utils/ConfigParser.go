package utils

import (
	"encoding/json"
	"os"

	"github.com/rideee/go-universal-health-check/src/models"
)

type ConfigParser struct {
	ServersConfigFile string
	ServersConfigData models.ServersConfig
	ServerTypesMap    map[string]models.ServerType
}

// Constructor.
func NewConfigParser(serversConfigFile string) *ConfigParser {
	cp := new(ConfigParser)
	cp.ServersConfigFile = serversConfigFile

	// Import json file.
	serversFileData, err := os.ReadFile(serversConfigFile)
	if err != nil {
		panic(err)
	}

	// Unmarshal json.
	err = json.Unmarshal(serversFileData, &cp.ServersConfigData)
	if err != nil {
		panic(err)
	}

	// Make a server types object map[string]models.ServerType.
	srvTypesMap := make(map[string]models.ServerType)
	for _, srvType := range cp.ServersConfigData.ServerTypes {
		srvTypesMap[srvType.Name] = srvType
	}
	cp.ServerTypesMap = srvTypesMap

	return cp
}
