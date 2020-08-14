package config

import (
	"github.com/go-playground/assert/v2"
	"github.com/goudai-projects/gdops/log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	log.Infof("Addr: %s, Port: %d", config.Server.Address, config.Server.Port)
	assert.Equal(t, config.Server.Address, "0.0.0.0")
	assert.Equal(t, config.Server.Port, 8080)
}
