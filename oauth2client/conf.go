package oauth2client

import (
	"encoding/json"
	"io/ioutil"
)

// Config is ...
type Config struct {
	ClientConfig ClientConfig `json:"client_config"`
}

// ClientConfig is ...
type ClientConfig struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Endpoint     Endpoint `json:"endpoint"`
	RedirectURL  string   `json:"redirect_uri"`
	Scopes       []string `json:"scopes"`
}

// Endpoint is ...
type Endpoint struct {
	AuthURL  string `json:"auth_url"`
	TokenURL string `json:"token_url"`
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
