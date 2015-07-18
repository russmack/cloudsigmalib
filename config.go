package cloudsigmalib

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a struct to hold all config data.
type Config struct {
	login Login
}

// Login is a struct to hold login config data.
type Login struct {
	Username string
	Password string
}

// NewConfig returns a new, empty Config.
func NewConfig() *Config {
	return &Config{}
}

// Load retrieves config data and sets the properties of the Config object.
func (c *Config) Load() (*Config, error) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return c, err
	}
	login := &Login{}
	err = json.Unmarshal(f, login)
	if err != nil {
		return c, err
	}
	c.login = *login
	return c, err
}

// Login is the getter of the Config.login field.
func (c *Config) Login() Login {
	return c.login
}
