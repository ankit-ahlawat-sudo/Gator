package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configData, err := os.ReadFile(fullPath)
	if err != nil && !os.IsNotExist(err) {
		return Config{}, err
	}

	var c Config

	if err := json.Unmarshal(configData, &c); err!= nil {
		return Config{}, err
	}

	return c, nil

}

func getConfigFilePath() (string, error) {
	homeLocation, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeLocation, configFileName)
	return fullPath, nil
}

func write(cfg Config) (error) {
	fullPath, err := getConfigFilePath()
		if err != nil {
			return err
		}

		data, err := json.Marshal(cfg)
		if err != nil {
			return err
		}
		if err := os.WriteFile(fullPath, data, 0644); err != nil {
			return err
		}
		

		return nil
}
