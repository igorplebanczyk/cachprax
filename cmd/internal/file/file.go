package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ServerInfo struct {
	PID    int    `json:"pid"`
	Origin string `json:"origin"`
	Port   int    `json:"port"`
}

func SaveDataToFile(pid int, origin string, port int) error {
	serverInfo := ServerInfo{
		PID:    pid,
		Origin: origin,
		Port:   port,
	}

	serverInfoJSON, err := json.Marshal(serverInfo)
	if err != nil {
		return fmt.Errorf("could not marshal server info to JSON: %v", err)
	}

	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	err = os.WriteFile(pidFile, serverInfoJSON, 0644)
	if err != nil {
		return fmt.Errorf("could not write server info to file: %v", err)
	}

	return nil
}

func GetDataFromFile() (ServerInfo, error) {
	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	fileContent, err := os.ReadFile(pidFile)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("could not read JSON file: %v", err)
	}

	serverInfo := ServerInfo{}
	err = json.Unmarshal(fileContent, &serverInfo)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return serverInfo, nil
}
