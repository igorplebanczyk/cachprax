package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
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
		return fmt.Errorf("could not write server info to state: %v", err)
	}

	return nil
}

func GetDataFromFile() (ServerInfo, error) {
	pidFile := filepath.Join(os.TempDir(), "cachprax.json")
	fileContent, err := os.ReadFile(pidFile)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("could not read JSON state: %v", err)
	}

	serverInfo := ServerInfo{}
	err = json.Unmarshal(fileContent, &serverInfo)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return serverInfo, nil
}

func IsProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(os.Signal(syscall.Signal(0)))
	return err == nil
}
