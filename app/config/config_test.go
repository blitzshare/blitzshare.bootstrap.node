package config_test

import (
	"os"
	"testing"

	"github.com/blitzshare/blitzshare.bootstrap.node/app/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	setUp()
	cfg, err := config.Load()

	assert.Nil(t, err, "Unable to log the config")
	assert.Equal(t, "38643", cfg.Server.Port)
	assert.Equal(t, "0.0.0.0", cfg.Server.Host)

	tearDown()
}

func setUp() {
	_ = os.Setenv("PORT", "38643")
	_ = os.Setenv("HOST", "0.0.0.0")
}

func tearDown() {
	_ = os.Unsetenv("PORT")
	_ = os.Unsetenv("ENV")
}
