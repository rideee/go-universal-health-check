package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/rideee/go-universal-health-check/src/models"
)

type ConfigParser struct {
	ConfigDir         string
	GeneralConfigFile string
	ServersConfigFile string
	GeneralConfigData models.GeneralConfig
	ServersConfigData models.ServersConfig
	ServerTypesMap    map[string]models.ServerType
}

// Constructor.
func NewConfigParser(configDir string) *ConfigParser {

	cp := new(ConfigParser)
	cp.ConfigDir = configDir
	cp.GeneralConfigFile = filepath.Join(cp.ConfigDir, "general.json")
	cp.ServersConfigFile = filepath.Join(cp.ConfigDir, "servers.json")

	// Import json files.
	generalFileData, err := os.ReadFile(cp.GeneralConfigFile)
	if err != nil {
		log.Panicf("Something went wrong while reading the general.json file.\n\n%v", err)
	}

	serversFileData, err := os.ReadFile(cp.ServersConfigFile)
	if err != nil {
		log.Panicf("Something went wrong while reading the servers.json file.\n\n%v", err)
	}

	// Unmarshal json.
	err = json.Unmarshal(generalFileData, &cp.GeneralConfigData)
	if err != nil {
		panic(err)
	}

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
