package oauth2client

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
)

// Config is ...
type Config struct {
	ClientConfig ClientConfig `json:"client_config"`
}

// ClientConfig is ...
type ClientConfig struct {
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"client_secret"`
	Endpoint     oauth2.Endpoint `json:"endpoint"`
	RedirectURL  string          `json:"redirect_uri"`
	Scopes       []string        `json:"scopes"`
}

// NewConfig is ...
func NewConfig(configFilePath string) (*Config, error) {
	config := new(Config)
	err := config.readConfig(configFilePath)
	return config, err
}

func (config *Config) readConfig(configFilePath string) error {
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	return nil
}
