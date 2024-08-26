package ev3lib

import (
	"fmt"
	"os"
)

type ConfigNotFoundError struct{ name string }

func (c ConfigNotFoundError) Error() string {
	return fmt.Sprintf("could not find config for %v", c.name)
}

var ConfigManager *configManager = &configManager{registeredConfigs: map[string]MenuConfig{}}

type configManager struct {
	registeredConfigs map[string]MenuConfig
}

func (c *configManager) Register(name string, config MenuConfig) {
	c.registeredConfigs[name] = config
}

func (c *configManager) GetConfig() (MenuConfig, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	config, found := c.registeredConfigs[hostname]
	if !found {
		return nil, ConfigNotFoundError{name: hostname}
	}

	return config, nil
}
